import React from 'react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';

export const SimulationPage: React.FC = () => {
  return (
    <div className="space-y-6">
      <div>
        <h2 className="text-3xl font-bold">模擬トレード</h2>
        <p className="text-muted-foreground mt-2">
          リスクなしで実践的なトレードを体験しましょう
        </p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <Card>
          <CardHeader>
            <CardTitle>トレードシミュレーション</CardTitle>
            <CardDescription>仮想資金でトレードを実行</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="bg-gray-100 h-64 rounded-lg flex items-center justify-center mb-4">
              <p className="text-muted-foreground">
                トレードチャートがここに表示されます
              </p>
            </div>
            <div className="grid grid-cols-2 gap-4">
              <Button variant="default">買い注文</Button>
              <Button variant="destructive">売り注文</Button>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>保有ポジション</CardTitle>
            <CardDescription>現在の仮想ポートフォリオ</CardDescription>
          </CardHeader>
          <CardContent>
            <p className="text-muted-foreground">
              保有銘柄の一覧が表示されます
            </p>
          </CardContent>
        </Card>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>AIトレード日記</CardTitle>
          <CardDescription>あなたのトレードをAIが分析</CardDescription>
        </CardHeader>
        <CardContent>
          <p className="text-muted-foreground">
            過去のトレード履歴とAIによる分析が表示されます
          </p>
        </CardContent>
      </Card>
    </div>
  );
};
