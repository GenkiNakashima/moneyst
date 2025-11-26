import { useState, useEffect } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { communityApi, type Post } from '@/api/community';
import { Search, TrendingUp, MessageSquare } from 'lucide-react';

export function CommunityPage() {
  const [newPostContent, setNewPostContent] = useState('');
  const [searchQuery, setSearchQuery] = useState('');
  const [isSearching, setIsSearching] = useState(false);
  const queryClient = useQueryClient();

  // Fetch posts
  const { data: posts, isLoading } = useQuery({
    queryKey: ['community-posts'],
    queryFn: () => communityApi.getPosts(),
  });

  // Fetch trending posts
  const { data: trendingPosts } = useQuery({
    queryKey: ['trending-posts'],
    queryFn: () => communityApi.getTrendingPosts(),
  });

  // Search posts
  const { data: searchResults } = useQuery({
    queryKey: ['search-posts', searchQuery],
    queryFn: () => communityApi.searchPosts(searchQuery),
    enabled: isSearching && searchQuery.length > 0,
  });

  // Create post mutation
  const createPostMutation = useMutation({
    mutationFn: communityApi.createPost,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['community-posts'] });
      setNewPostContent('');
    },
  });

  // Toggle like mutation
  const toggleLikeMutation = useMutation({
    mutationFn: communityApi.toggleLike,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['community-posts'] });
      queryClient.invalidateQueries({ queryKey: ['trending-posts'] });
    },
  });

  const handleCreatePost = (e: React.FormEvent) => {
    e.preventDefault();
    if (newPostContent.trim()) {
      createPostMutation.mutate({ content: newPostContent });
    }
  };

  const handleSearch = () => {
    if (searchQuery.trim()) {
      setIsSearching(true);
    } else {
      setIsSearching(false);
    }
  };

  const displayedPosts = isSearching && searchQuery ? searchResults : posts;

  return (
    <div className="max-w-4xl mx-auto p-6 space-y-6">
      <div className="flex items-center justify-between">
        <h1 className="text-3xl font-bold">コミュニティ</h1>
        <MessageSquare className="w-8 h-8" />
      </div>

      {/* Trending Posts */}
      {trendingPosts && trendingPosts.length > 0 && (
        <div className="border rounded-lg p-4 bg-gradient-to-r from-orange-50 to-yellow-50">
          <div className="flex items-center gap-2 mb-3">
            <TrendingUp className="w-5 h-5 text-orange-600" />
            <h2 className="text-lg font-semibold">急上昇中の投稿 TOP3</h2>
          </div>
          <div className="space-y-2">
            {trendingPosts.map((post, index) => (
              <div key={post.id} className="flex gap-3 p-3 bg-white rounded border">
                <div className="text-2xl font-bold text-orange-600">#{index + 1}</div>
                <div className="flex-1">
                  <p className="font-medium">{post.user?.username}</p>
                  <p className="text-sm text-gray-600 line-clamp-2">{post.content}</p>
                  <div className="flex gap-4 text-xs text-gray-500 mt-2">
                    <span>🔥 {post.likes_count}</span>
                    <span>💬 {post.replies_count}</span>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>
      )}

      {/* Search Bar */}
      <div className="flex gap-2">
        <input
          type="text"
          placeholder="投稿を検索..."
          className="flex-1 px-4 py-2 border rounded-lg"
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          onKeyPress={(e) => e.key === 'Enter' && handleSearch()}
        />
        <button
          onClick={handleSearch}
          className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 flex items-center gap-2"
        >
          <Search className="w-4 h-4" />
          検索
        </button>
        {isSearching && (
          <button
            onClick={() => {
              setIsSearching(false);
              setSearchQuery('');
            }}
            className="px-4 py-2 border rounded-lg hover:bg-gray-50"
          >
            クリア
          </button>
        )}
      </div>

      {/* Create Post Form */}
      <form onSubmit={handleCreatePost} className="border rounded-lg p-4 bg-white">
        <textarea
          className="w-full px-3 py-2 border rounded-lg resize-none"
          rows={4}
          placeholder="投資に関する情報を共有しましょう...&#10;&#10;💡 ヒント: @checkAI とメンションするとAIが質問に答えます"
          value={newPostContent}
          onChange={(e) => setNewPostContent(e.target.value)}
        />
        <div className="flex justify-between items-center mt-3">
          <span className="text-sm text-gray-500">
            {newPostContent.length}/5000
          </span>
          <button
            type="submit"
            disabled={!newPostContent.trim() || createPostMutation.isPending}
            className="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:bg-gray-400 disabled:cursor-not-allowed"
          >
            {createPostMutation.isPending ? '投稿中...' : '投稿する'}
          </button>
        </div>
        <p className="text-xs text-gray-500 mt-2">
          ⚠️ 投稿は自動的にAIファクトチェックされます
        </p>
      </form>

      {/* Posts List */}
      <div className="space-y-4">
        {isLoading && <div className="text-center py-8">読み込み中...</div>}

        {displayedPosts?.map((post) => (
          <PostCard
            key={post.id}
            post={post}
            onLike={() => toggleLikeMutation.mutate(post.id)}
          />
        ))}

        {displayedPosts?.length === 0 && (
          <div className="text-center py-8 text-gray-500">
            {isSearching ? '検索結果が見つかりませんでした' : '投稿がまだありません'}
          </div>
        )}
      </div>
    </div>
  );
}

// PostCard Component
function PostCard({ post, onLike }: { post: Post; onLike: () => void }) {
  const [showReplies, setShowReplies] = useState(false);
  const [replyContent, setReplyContent] = useState('');
  const queryClient = useQueryClient();

  const { data: postDetails } = useQuery({
    queryKey: ['post-details', post.id],
    queryFn: () => communityApi.getPost(post.id),
    enabled: showReplies,
  });

  const createReplyMutation = useMutation({
    mutationFn: (content: string) =>
      communityApi.createReply(post.id, { content }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['post-details', post.id] });
      queryClient.invalidateQueries({ queryKey: ['community-posts'] });
      setReplyContent('');
    },
  });

  const handleReply = (e: React.FormEvent) => {
    e.preventDefault();
    if (replyContent.trim()) {
      createReplyMutation.mutate(replyContent);
    }
  };

  return (
    <div className="border rounded-lg p-4 bg-white">
      {/* Post Header */}
      <div className="flex items-start justify-between mb-3">
        <div>
          <p className="font-semibold">{post.user?.username}</p>
          <p className="text-xs text-gray-500">
            {new Date(post.created_at).toLocaleString('ja-JP')}
          </p>
        </div>
        {post.fact_check_status === 'flagged' && (
          <span className="text-xs px-2 py-1 bg-yellow-100 text-yellow-800 rounded">
            ⚠️ 要確認
          </span>
        )}
        {post.fact_check_status === 'approved' && (
          <span className="text-xs px-2 py-1 bg-green-100 text-green-800 rounded">
            ✓ 確認済み
          </span>
        )}
      </div>

      {/* Post Content */}
      <p className="text-gray-800 whitespace-pre-wrap mb-3">{post.content}</p>

      {/* Fact Check Result */}
      {post.fact_check_result && post.fact_check_status === 'flagged' && (
        <div className="mb-3 p-3 bg-yellow-50 border border-yellow-200 rounded text-sm">
          <p className="text-yellow-800">{post.fact_check_result}</p>
        </div>
      )}

      {/* Post Actions */}
      <div className="flex gap-4 pt-3 border-t">
        <button
          onClick={onLike}
          className={`flex items-center gap-1 text-sm ${
            post.user_liked ? 'text-red-600' : 'text-gray-600 hover:text-red-600'
          }`}
        >
          <span className={post.user_liked ? '🔥' : '🔥'}>
            {post.user_liked ? '🔥' : '🔥'}
          </span>
          {post.likes_count}
        </button>
        <button
          onClick={() => setShowReplies(!showReplies)}
          className="flex items-center gap-1 text-sm text-gray-600 hover:text-blue-600"
        >
          💬 {post.replies_count}
        </button>
      </div>

      {/* Replies Section */}
      {showReplies && (
        <div className="mt-4 pt-4 border-t space-y-3">
          {postDetails?.replies.map((reply) => (
            <div
              key={reply.id}
              className={`p-3 rounded ${
                reply.is_ai_response ? 'bg-blue-50 border border-blue-200' : 'bg-gray-50'
              }`}
            >
              <div className="flex items-center gap-2 mb-1">
                {reply.is_ai_response && <span className="text-xs">🤖 AI</span>}
                <p className="text-sm font-medium">{reply.user?.username}</p>
                <p className="text-xs text-gray-500">
                  {new Date(reply.created_at).toLocaleString('ja-JP')}
                </p>
              </div>
              <p className="text-sm whitespace-pre-wrap">{reply.content}</p>
            </div>
          ))}

          {/* Reply Form */}
          <form onSubmit={handleReply} className="flex gap-2">
            <input
              type="text"
              placeholder="返信を入力... (@checkAI で AI に質問)"
              className="flex-1 px-3 py-2 text-sm border rounded"
              value={replyContent}
              onChange={(e) => setReplyContent(e.target.value)}
            />
            <button
              type="submit"
              disabled={!replyContent.trim() || createReplyMutation.isPending}
              className="px-4 py-2 text-sm bg-blue-600 text-white rounded hover:bg-blue-700 disabled:bg-gray-400"
            >
              返信
            </button>
          </form>
        </div>
      )}
    </div>
  );
}
