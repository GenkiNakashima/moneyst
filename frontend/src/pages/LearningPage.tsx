import React from 'react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';

export const LearningPage: React.FC = () => {
  return (
    <div className="space-y-6">
      <div>
        <h2 className="text-3xl font-bold">インタラクティブ学習</h2>
        <p className="text-muted-foreground mt-2">
          AIコーチと一緒にチャート分析を学びましょう
        </p>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>チャート学習</CardTitle>
          <CardDescription>
            実際のチャートをクリックして、AIコーチに質問しましょう
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="bg-gray-100 h-96 rounded-lg flex items-center justify-center">
            <p className="text-muted-foreground">
              チャートコンポーネント（TradingView Lightweight Charts）がここに表示されます
            </p>
          </div>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>パーソナルカリキュラム</CardTitle>
        </CardHeader>
        <CardContent>
          <p className="text-muted-foreground">
            あなた専用の学習カリキュラムが表示されます
          </p>
        </CardContent>
      </Card>
    </div>
  );
};
