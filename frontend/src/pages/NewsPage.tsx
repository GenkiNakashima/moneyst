import React from 'react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';

export const NewsPage: React.FC = () => {
  return (
    <div className="space-y-6">
      <div>
        <h2 className="text-3xl font-bold">パーソナライズドニュース</h2>
        <p className="text-muted-foreground mt-2">
          あなたのレベルに合わせてAIが要約したニュース
        </p>
      </div>

      <div className="space-y-4">
        {[1, 2, 3, 4, 5].map((i) => (
          <Card key={i}>
            <CardHeader>
              <CardTitle className="text-lg">ニュース記事タイトル {i}</CardTitle>
              <CardDescription>2024-11-18</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                <p className="text-sm">
                  AI要約: ここにニュースの要約が表示されます。
                  ユーザーの理解度に合わせて難易度が調整されています。
                </p>
                <div className="bg-blue-50 p-4 rounded-lg">
                  <p className="text-sm font-medium text-blue-900">AI洞察</p>
                  <p className="text-sm text-blue-700 mt-1">
                    このニュースは市場にポジティブな影響を与える可能性があります。
                  </p>
                </div>
              </div>
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  );
};
