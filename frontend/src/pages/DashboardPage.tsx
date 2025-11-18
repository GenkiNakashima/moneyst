import React, { useEffect, useState } from 'react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { useAuth } from '@/contexts/AuthContext';
import { apiClient } from '@/api/client';

interface DailyLesson {
  id: string;
  content: string;
  completed: boolean;
  points_earned: number;
}

export const DashboardPage: React.FC = () => {
  const { user } = useAuth();
  const [dailyLesson, setDailyLesson] = useState<DailyLesson | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchDailyLesson = async () => {
      try {
        const response = await apiClient.get('/api/lessons/daily');
        setDailyLesson(response.data);
      } catch (error) {
        console.error('Failed to fetch daily lesson:', error);
      } finally {
        setIsLoading(false);
      }
    };

    fetchDailyLesson();
  }, []);

  const handleCompleteLesson = async () => {
    if (!dailyLesson) return;

    try {
      const response = await apiClient.post(`/api/lessons/${dailyLesson.id}/complete`);
      setDailyLesson(response.data);
    } catch (error) {
      console.error('Failed to complete lesson:', error);
    }
  };

  return (
    <div className="space-y-6">
      {/* Welcome Header */}
      <div>
        <h2 className="text-3xl font-bold">こんにちは、{user?.username}さん</h2>
        <p className="text-muted-foreground mt-2">
          今日も一緒に投資の学習を進めましょう
        </p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        {/* Daily Lesson Card */}
        <Card>
          <CardHeader>
            <CardTitle>今日の3分レッスン</CardTitle>
            <CardDescription>毎日少しずつ、着実にスキルアップ</CardDescription>
          </CardHeader>
          <CardContent>
            {isLoading ? (
              <p>読み込み中...</p>
            ) : dailyLesson ? (
              <div className="space-y-4">
                <div className="prose prose-sm max-w-none">
                  <div dangerouslySetInnerHTML={{ __html: dailyLesson.content.replace(/\n/g, '<br />') }} />
                </div>
                {!dailyLesson.completed ? (
                  <Button onClick={handleCompleteLesson} className="w-full">
                    レッスンを完了する
                  </Button>
                ) : (
                  <div className="text-center py-4 bg-green-50 rounded-lg">
                    <p className="text-green-700 font-medium">
                      完了済み！ {dailyLesson.points_earned} ポイント獲得
                    </p>
                  </div>
                )}
              </div>
            ) : (
              <p>レッスンが見つかりません</p>
            )}
          </CardContent>
        </Card>

        {/* User Status Card */}
        <Card>
          <CardHeader>
            <CardTitle>あなたの分析結果</CardTitle>
            <CardDescription>AIがあなたの投資スタイルを分析しました</CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div>
              <p className="text-sm font-medium text-muted-foreground">リスク許容度</p>
              <p className="text-lg font-semibold">
                {user?.risk_tolerance === 'low' ? '保守的' :
                 user?.risk_tolerance === 'high' ? '積極的' :
                 user?.risk_tolerance === 'medium' ? '中立' : '未設定'}
              </p>
            </div>
            <div>
              <p className="text-sm font-medium text-muted-foreground">学習スタイル</p>
              <p className="text-lg font-semibold">
                {user?.learning_style === 'text' ? 'テキスト型' :
                 user?.learning_style === 'visual' ? '図解型' :
                 user?.learning_style === 'conversational' ? '会話型' : '未設定'}
              </p>
            </div>
            <div>
              <p className="text-sm font-medium text-muted-foreground">投資経験</p>
              <p className="text-lg font-semibold">
                {user?.investment_experience === 'none' ? '未経験' :
                 user?.investment_experience === 'beginner' ? '初心者' :
                 user?.investment_experience === 'intermediate' ? '中級者' : '未設定'}
              </p>
            </div>
          </CardContent>
        </Card>

        {/* Quick Stats */}
        <Card className="md:col-span-2">
          <CardHeader>
            <CardTitle>最近のアクティビティ</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-muted-foreground">
              ここに最近の学習履歴や模擬トレードの成績が表示されます
            </p>
          </CardContent>
        </Card>
      </div>
    </div>
  );
};
