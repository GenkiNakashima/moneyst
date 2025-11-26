import { apiClient } from './client';

export interface User {
  id: string;
  email: string;
  username: string;
}

export interface Post {
  id: string;
  user_id: string;
  user?: User;
  content: string;
  is_fact_checked: boolean;
  fact_check_result?: string;
  fact_check_status: 'pending' | 'approved' | 'flagged' | 'removed';
  likes_count: number;
  replies_count: number;
  user_liked: boolean;
  created_at: string;
  updated_at: string;
}

export interface Reply {
  id: string;
  post_id: string;
  user_id: string;
  user?: User;
  content: string;
  is_ai_response: boolean;
  created_at: string;
  updated_at: string;
}

export interface PostWithReplies {
  post: Post;
  replies: Reply[];
}

export interface CreatePostRequest {
  content: string;
}

export interface CreateReplyRequest {
  content: string;
}

export const communityApi = {
  // Get all posts (paginated)
  getPosts: async (limit = 20, offset = 0): Promise<Post[]> => {
    const response = await apiClient.get('/api/community/posts', {
      params: { limit, offset },
    });
    return response.data;
  },

  // Get a single post with replies
  getPost: async (id: string): Promise<PostWithReplies> => {
    const response = await apiClient.get(`/api/community/posts/${id}`);
    return response.data;
  },

  // Create a new post
  createPost: async (data: CreatePostRequest): Promise<Post> => {
    const response = await apiClient.post('/api/community/posts', data);
    return response.data;
  },

  // Delete a post
  deletePost: async (id: string): Promise<void> => {
    await apiClient.delete(`/api/community/posts/${id}`);
  },

  // Create a reply to a post
  createReply: async (postId: string, data: CreateReplyRequest): Promise<Reply> => {
    const response = await apiClient.post(`/api/community/posts/${postId}/replies`, data);
    return response.data;
  },

  // Toggle like on a post
  toggleLike: async (postId: string): Promise<{ liked: boolean; message: string }> => {
    const response = await apiClient.post(`/api/community/posts/${postId}/like`);
    return response.data;
  },

  // Search posts
  searchPosts: async (query: string, limit = 20, offset = 0): Promise<Post[]> => {
    const response = await apiClient.get('/api/community/posts/search', {
      params: { q: query, limit, offset },
    });
    return response.data;
  },

  // Get trending posts (top 3)
  getTrendingPosts: async (): Promise<Post[]> => {
    const response = await apiClient.get('/api/community/posts/trending');
    return response.data;
  },
};
