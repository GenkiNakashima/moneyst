import React from 'react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';

export const RiskDashboardPage: React.FC = () => {
  return (
    <div className="space-y-6">
      <div>
        <h2 className="text-3xl font-bold">リスク分析ダッシュボード</h2>
        <p className="text-muted-foreground mt-2">
          暗号資産のリスクを分かりやすく可視化
        </p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <Card>
          <CardHeader>
            <CardTitle>ボラティリティ指数</CardTitle>
            <CardDescription>価格変動の激しさ</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="bg-gray-100 h-64 rounded-lg flex items-center justify-center">
              <p className="text-muted-foreground">
                ボラティリティチャートがここに表示されます
              </p>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>リスク・リターン分析</CardTitle>
            <CardDescription>シャープレシオ</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="bg-gray-100 h-64 rounded-lg flex items-center justify-center">
              <p className="text-muted-foreground">
                リスク・リターン散布図がここに表示されます
              </p>
            </div>
          </CardContent>
        </Card>

        <Card className="md:col-span-2">
          <CardHeader>
            <CardTitle>暴落シミュレーション</CardTitle>
            <CardDescription>過去の暴落時の影響</CardDescription>
          </CardHeader>
          <CardContent>
            <p className="text-muted-foreground">
              各銘柄の過去の暴落時のパフォーマンスが表示されます
            </p>
          </CardContent>
        </Card>

        <Card className="md:col-span-2">
          <CardHeader>
            <CardTitle>AIパーソナルアラート</CardTitle>
          </CardHeader>
          <CardContent className="space-y-2">
            <div className="bg-yellow-50 p-4 rounded-lg">
              <p className="text-sm font-medium text-yellow-900">
                あなたのリスク許容度: 保守的
              </p>
              <p className="text-sm text-yellow-700 mt-1">
                この銘柄はあなたのリスク許容度に対して高すぎる可能性があります。
                慎重な検討をお勧めします。
              </p>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
};
