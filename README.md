# プロジェクト名: AI投資学習コンパニオン (仮)

## 1. 概要

本プロジェクトは、AIを活用してユーザー個別に最適化された投資学習体験を提供するWebアプリケーションを開発することを目的とする。従来の画一的な教材提供型アプリとは異なり、AIがユーザーのリスク許容度、学習スタイル、脳タイプを分析し、「伴走者」として継続的な学習と実践（模擬トレード）をサポートする。

最終目標は、投資初心者が挫折することなく、実践的な知識と「自分なりの投資の癖」を理解し、成長できるプラットフォームを構築することである。

## 2. 要求仕様（主要機能）

ユーザーが提示した7つの主要な実装ポイントに基づき、以下の機能を要求仕様として定義する。

### ✅ 1. AIによるパーソナルカリキュラム生成

**機能:**
* ユーザー登録時に「性格診断」「投資行動テスト」を実施する。
* アプリ内での「模擬トレード」（数回）の結果に基づき、AIが以下の要素を推定・判定する。
    * リスク許容度（積極的、中立、保守的）
    * 学習スタイル（文章型、図解型、会話型）
    * 投資経験の深さ（未経験、初心者、中級者）
    * 取引における「迷いやすいパターン」（例：損切りが遅い、利確が早い）
* 判定結果に基づき、ユーザー専用の「パーソナルカリキュラム」を自動生成する。
* カリキュラムは、ユーザーの弱点や学習スタイルに合わせて、教材の順序や提示方法（テキスト、図解、AI対話）を最適化する。

**目的:**
* 画一的な学習による「難しすぎる」「つまらない」といった挫折ポイントを排除し、学習継続率を向上させる。

### ✅ 2. AIコーチ付き インタラクティブ・チャート学習

**機能:**
* 実際の（または過去の）株価・暗号資産チャートを表示する。
* ユーザーがチャート上の特定の箇所をタップ（クリック）すると、AIコーチがポップアップやチャット形式で解説を提供する。
    * 例：「今この形は『ダブルボトム』の可能性があり、上昇転換のサインかもしれません」
    * 例：「初心者がこの場面でやりがちなミスは、焦って『飛びつき買い』することです」
* AIコーチの解説は、(✅ 1)で判定したユーザーの傾向（例：迷いやすいパターン）に応じてパーソナライズされる。
* （オプション）スマートフォンで撮影したチャート画像（紙や別モニター）を解析し、同様の解説を提供する。

**目的:**
* 投資初心者が最もつまずく「チャート読解」を、実践的な形式でサポートし、リアルタイムの疑問を解消する。

### ✅ 3. パーソナライズAI要約ニュース

**機能:**
* 外部API（News APIなど）から株・暗号資産関連のニュースをリアルタイムで取得する。
* AIがニュースコンテンツを解析・要約する。
* 要約の難易度は、(✅ 1)で判定したユーザーの理解度や投資経験に合わせて調整する（例：「小学生向け」「中級者向け」「専門家向け」）。
* AIは単なる要約に留まらず、以下の補足情報を生成・付与する。
    * 「このニュースが（関連する）株価や暗号資産にどう影響しそうか（ポジティブ/ネガティブ/中立）」
    * 「過去に類似のニュースが出た際、市場がどう動いたか」

**目的:**
* 情報過多なニュースを「自分ごと化」し、難解な経済ニュースを理解可能なレベルで提供することで、日々の情報収集を習慣化させる。

### ✅ 4. AI生成による「3分デイリーレッスン」

**機能:**
* 毎日1回、ユーザー専用の「3分でわかる今日の投資レッスン」をAIが自動生成し、ダッシュボードに提示する。
* レッスンの内容は、(✅ 1)のカリキュラムやユーザーの弱点（例：チャート問題、経済用語クイズ）に最適化される。
* レッスンを完了するとポイントやバッジが付与され、レベルアップするUI（ゲーミフィケーション要素）を実装する。

**目的:**
* 学習の習慣化（短時間でも毎日触れる）を促し、「着実に実力がつく感覚」を提供して継続率を高める。

### ✅ 5. AIフィードバック付き 模擬投資（シミュレーション）

**機能:**
* 過去の実際のチャートデータ（24時間/365日分）を使用した「模擬トレード」機能を提供する。
* ユーザーは仮想資金を用いて、任意のタイミングで売買を体験できる。
* トレード完了後（または一定期間後）、AIがその取引の判断を分析し、フィードバックを提供する。
    * 例：「エントリー（買い）のタイミングが少し早すぎたかもしれません」
    * 例：「損切りライン（-5%）の判断は適切でしたが、実行が遅れました」
    * 例：「この取引は、ニュースに強く反応する『モメンタム投資家』に近い動きです」
* (✅ 1)の診断結果（性格）と照らし合わせ、「あなたの性格タイプに近い投資家」を診断する。

**目的:**
* 座学だけでは得られない「実践感覚」を提供し、ゲーム感覚で投資判断を学ばせる。

### ✅ 6. AIによる「トレード日記」自動生成

**機能:**
* (✅ 5)の模擬トレード、または実トレード（API連携・オプション）の履歴をAIが読み取る。
* AIが以下の項目を含む「トレード日記」を自動生成する。
    * 取引日時、銘柄、売買理由（推定）
    * その時の主要ニュースや市場のセンチメント
    * AIによる評価：「良かった点」「改善が必要な点」
* ユーザーは生成された日記にコメントを追記できる。

**目的:**
* トレードの振り返り（日記作成）という面倒な作業をAIが代行することで、ユーザーが「自身の無意識の癖」に気づき、成長を加速させる。

### ✅ 7. 暗号資産リスク・ダッシュボード

**機能:**
* 暗号資産（主要銘柄）のリスクを可視化する専用ダッシュボードを提供する。
* 以下の指標を、専門用語を極力使わずに平易な言葉とグラフで説明する。
    * ボラティリティ指数（価格変動の激しさ）
    * リスク調整後リターン（シャープレシオなど）
    * 過去の暴落シミュレーション（例：コロナショック時、どの程度下落したか）
    * 外部リスク（ハッキング、規制ニュースなど）
* (✅ 1)のユーザーのリスク許容度に基づき、「あなたにとってこの銘柄のリスクは高すぎる（または適切）かもしれません」といったアラートを表示する。

**目的:**
* 暗号資産投資における「過剰なリスクテイク」を抑制し、初心者が自身の許容範囲内で資産形成を行えるようガイドする。

### ✅ 8. コミュニティ機能（情報共有）

**機能:**
* ユーザー同士が投資に関する情報を共有し、学び合えるコミュニティ機能を提供する。
* **投稿機能**
    * ユーザーは投資に関する情報、知識、体験談を自由に投稿できる。
    * すべての投稿は自動的にAIによるファクトチェックが実行される。
    * ファクトチェック結果は投稿に「✓ 確認済み」「⚠️ 要確認」として表示される。
    * 誤情報や不適切な内容はAIが検出し、警告を表示することで、コミュニティの質を維持する。
* **スレッド形式の返信**
    * 投稿に対してユーザーは返信（リプライ）をつけることができる。
    * 返信は投稿ごとにスレッド形式で表示され、議論を整理しやすくする。
* **AI統合（@checkAI機能）**
    * 投稿や返信で`@checkAI`とメンションすると、AIが質問に答えたり、投資に関する助言を提供する。
    * AI応答は「🤖 AI」アイコンで識別され、人間の投稿と明確に区別される。
    * 初心者がすぐに疑問を解消できるよう、24時間対応の「AIコーチ」として機能する。
* **検索機能**
    * 投稿を全文検索できる機能を提供し、過去の議論や情報を簡単に探せる。
    * PostgreSQLの全文検索インデックスを活用し、高速な検索を実現する。
* **トレンド機能（急上昇TOP3）**
    * 「いいね（🔥）」数と返信数が多い投稿をTOP3として表示する。
    * ユーザーは人気のある議論や話題をすぐに把握できる。
    * トレンドはリアルタイムで更新され、最新の関心事を可視化する。
* **いいね（Like）システム**
    * ユーザーは役立つ投稿に「🔥」を付けて評価できる。
    * いいね数は投稿の信頼性や人気度の指標となり、質の高い情報が自然に上位に表示される仕組み。

**目的:**
* 孤独な学習ではなく、コミュニティで学び合う環境を提供し、ユーザーのモチベーション向上と学習継続を促進する。
* AIファクトチェックにより、誤情報や詐欺的な投稿を抑制し、初心者が安心して情報を得られる環境を構築する。
* ユーザー同士の知識共有により、「実践的な投資ノウハウ」や「失敗談」など、教材では学べない生の情報を提供する。

## 3. データベース設計（主要テーブル）



### users (ユーザー)
* ユーザーの基本情報と、AIによる分析結果を格納する。

| カラム名 | 型 | 説明 |
| :--- | :--- | :--- |
| `id` | `uuid` (PK) | ユーザーID |
| `email` | `varchar` | メールアドレス (認証用) |
| `hashed_password` | `varchar` | ハッシュ化パスワード |
| `username` | `varchar` | 表示名 |
| `risk_tolerance` | `varchar` | リスク許容度 (low, medium, high) |
| `learning_style` | `varchar` | 学習スタイル (text, visual, conversational) |
| `investment_experience` | `varchar` | 投資経験 (none, beginner, intermediate) |
| `trading_pattern_memo` | `text` | AIによる行動パターン分析メモ |
| `created_at` | `timestamp` | 作成日時 |
| `updated_at` | `timestamp` | 更新日時 |

### curriculums (カリキュラム)
* AIが生成したユーザー個別の学習進捗を管理する。

| カラム名 | 型 | 説明 |
| :--- | :--- | :--- |
| `id` | `uuid` (PK) | カリキュラムID |
| `user_id` | `uuid` (FK) | ユーザーID |
| `title` | `varchar` | レッスンタイトル |
| `module_type` | `varchar` | モジュール種別 (chart, news, simulation, quiz) |
| `content_id` | `uuid` | 各種コンテンツID (lessons, quizzesなど) |
| `status` | `varchar` | 進捗 (pending, in_progress, completed) |
| `order` | `int` | 学習順序 |
| `created_at` | `timestamp` | 作成日時 |

### lessons (レッスン教材)
* 「3分デイリーレッスン」やチャート学習の元となる教材マスター。

| カラム名 | 型 | 説明 |
| :--- | :--- | :--- |
| `id` | `uuid` (PK) | レッスンID |
| `title` | `varchar` | タイトル |
| `content_body` | `text` | レッスン内容 (Markdown or JSON) |
| `difficulty_level` | `int` | 難易度 (1-5) |
| `category` | `varchar` | カテゴリ (chart, glossary, economy) |
| `estimated_time_min` | `int` | 推定学習時間（分） |

### simulation_trades (模擬トレード履歴)
* 模擬トレードの全履歴。トレード日記の元データ。

| カラム名 | 型 | 説明 |
| :--- | :--- | :--- |
| `id` | `uuid` (PK) | トレードID |
| `user_id` | `uuid` (FK) | ユーザーID |
| `symbol` | `varchar` | 銘柄コード (例: BTC/USD, AAPL) |
| `trade_type` | `varchar` | 売買 (buy, sell) |
| `entry_price` | `decimal` | 約定価格 |
| `quantity` | `decimal` | 数量 |
| `entry_at` | `timestamp` | 約定日時（シミュレーション内時刻） |
| `exit_price` | `decimal` | 決済価格 (オプション) |
| `exit_at` | `timestamp` | 決済日時 (オプション) |
| `profit_loss` | `decimal` | 損益 (オプション) |

### trade_diaries (トレード日記)
* AIが自動生成したトレード日記。

| カラム名 | 型 | 説明 |
| :--- | :--- | :--- |
| `id` | `uuid` (PK) | 日記ID |
| `user_id` | `uuid` (FK) | ユーザーID |
| `trade_id` | `uuid` (FK) | 関連するトレードID (simulation_trades.id) |
| `generated_analysis` | `text` | AIによる自動生成分析（良かった点、改善点） |
| `market_context` | `text` | AIによる市場状況の要約 |
| `user_memo` | `text` | ユーザーによる追記メモ |
| `diary_date` | `date` | 日記の日付 |
| `created_at` | `timestamp` | 作成日時 |

### summarized_news (要約ニュース)
* AIが要約し、ユーザーごとにパーソナライズしたニュース記事。

| カラム名 | 型 | 説明 |
| :--- | :--- | :--- |
| `id` | `uuid` (PK) | ニュースID |
| `user_id` | `uuid` (FK) | ユーザーID |
| `original_url` | `varchar` | 元記事URL |
| `original_title` | `varchar` | 元記事タイトル |
| `summarized_title` | `varchar` | AI要約タイトル |
| `summarized_content` | `text` | AI要約本文（難易度調整済み） |
| `ai_insight` | `text` | AIによる洞察（市場への影響予測など） |
| `published_at` | `timestamp` | 配信日時 |
| `read_status` | `boolean` | 既読フラグ |

### posts (投稿)
* コミュニティ機能におけるユーザー投稿。

| カラム名 | 型 | 説明 |
| :--- | :--- | :--- |
| `id` | `uuid` (PK) | 投稿ID |
| `user_id` | `uuid` (FK) | 投稿者のユーザーID |
| `content` | `text` | 投稿内容 |
| `is_fact_checked` | `boolean` | ファクトチェック済みフラグ |
| `fact_check_result` | `text` | AIファクトチェック結果（警告メッセージなど） |
| `fact_check_status` | `varchar` | ステータス (pending, approved, flagged, removed) |
| `created_at` | `timestamp` | 作成日時 |
| `updated_at` | `timestamp` | 更新日時 |

### replies (返信)
* 投稿に対するスレッド形式の返信。

| カラム名 | 型 | 説明 |
| :--- | :--- | :--- |
| `id` | `uuid` (PK) | 返信ID |
| `post_id` | `uuid` (FK) | 関連する投稿ID (posts.id) |
| `user_id` | `uuid` (FK) | 返信者のユーザーID |
| `content` | `text` | 返信内容 |
| `is_ai_response` | `boolean` | AI応答フラグ（@checkAI応答の場合true） |
| `created_at` | `timestamp` | 作成日時 |
| `updated_at` | `timestamp` | 更新日時 |

### post_likes (投稿のいいね)
* 投稿に対するいいね（トレンド機能のデータ）。

| カラム名 | 型 | 説明 |
| :--- | :--- | :--- |
| `id` | `uuid` (PK) | いいねID |
| `post_id` | `uuid` (FK) | 関連する投稿ID (posts.id) |
| `user_id` | `uuid` (FK) | いいねしたユーザーID |
| `created_at` | `timestamp` | 作成日時 |

---

## 4. 技術構成案 (Technical Architecture)

本プロジェクトの技術スタックおよびアーキテクチャは、スケーラビリティ、リアルタイム性、AI連携の容易性を考慮し、以下の通り提案する。

### 4.1. システムアーキテクチャ概要



* **フロントエンド:** React (TypeScript) で構築されたSPA。静的ファイルはAWS S3にホスティングされ、CloudFrontを通じて配信される。
* **バックエンド:** Goで構築されたAPIサーバー。Dockerコンテナ化し、AWS ECS (Fargate) 上で実行される。
* **データベース:** Amazon RDS for PostgreSQL を使用し、データの永続化を行う。
* **AI・非同期処理:**
    * AIによる分析・生成（カリキュラム生成、ニュース要約、日記生成など）は、計算負荷が高いため、APIサーバーから直接実行せず、非同期処理とする。
    * APIサーバーはリクエストを受け付けたら、SQSなどのキューにタスクを投入。
    * AWS Lambda（またはECSの別タスク）がキューを監視し、タスクを実行。AIモデル（Gemini API, OpenAI APIなど）と通信する。
    * 処理結果はDBに書き込まれ、必要に応じてWebSocketやSSEを通じてフロントエンドに通知される。

### 4.2. フロントエンド (Frontend)

* **フレームワーク:** React 18+
* **言語:** TypeScript
* **状態管理:** Zustand または React Context。(グローバルな状態が複雑になりすぎない限り、軽量なものを選択)
* **UIコンポーネント:** shadcn/ui または Material-UI (MUI)。迅速な開発とアクセシビリティを両立する。
* **データ取得:** TanStack Query (React Query)。APIのキャッシュ管理、非同期状態の管理を簡素化する。
* **チャート:** TradingView (Lightweight Charts) または ECharts。インタラクティブなチャート表示（タップイベント取得）に対応可能なライブラリを選定する。
* **ビルド:** Vite

### 4.3. バックエンド (Backend)

* **言語:** Go (Golang) 1.21+
* **フレームワーク:** Gin。高速なルーティングとシンプルなAPI開発に適している。
* **ORM:** GORM。PostgreSQLとの連携を効率化する。
* **認証:** JWT (JSON Web Tokens) によるセッション管理。

### 4.4. データベース (Database)

* **RDBMS:** PostgreSQL 15+
* **ホスティング:** Amazon RDS (Relational Database Service)。スケーラビリティ、バックアップ、可用性をAWSに委任する。

### 4.5. インフラストラクチャ (Infrastructure)

* **コンテナ:** Docker
* **コンテナオーケストレーション:** AWS ECS (Elastic Container Service) on Fargate。サーバーレスのコンテナ実行環境により、インフラ管理コストを削減する。
* **CI/CD:** GitHub Actions または AWS CodePipeline。Gitリポジトリへのプッシュをトリガーに、自動でテスト、ビルド、Dockerイメージのプッシュ、ECSへのデプロイを実行する。
* **ストレージ:** AWS S3。フロントエンドの静的ファイル、ユーザーがアップロードした画像（チャート写真など）の保存場所。
* **CDN:** AWS CloudFront。S3の静的ファイルとAPIへのアクセスを高速化・キャッシュする。

### 4.6. 主要APIエンドポイント (抜粋)

* **認証**
    * `POST /auth/register` : ユーザー登録
    * `POST /auth/login` : ログイン (JWT発行)
    * `GET /auth/me` : ユーザー情報取得 (認証確認)
* **ユーザー分析**
    * `POST /users/me/analyze` : ユーザー分析テスト（性格診断など）の結果を送信
    * `GET /users/me/profile` : AIによる分析結果（リスク許容度など）を取得
* **学習**
    * `GET /curriculums/me` : パーソナルカリキュラムを取得
    * `POST /curriculums/me/complete` : レッスンの完了をマーク
    * `GET /lessons/daily` : 今日の3分レッスンを取得
* **AIコーチ（チャート）**
    * `POST /ai/coach/chart-insight` : チャートの特定箇所（座標、価格帯）を送信し、AIの解説を要求
* **ニュース**
    * `GET /news/personalized` : AI要約ニュースフィードを取得
* **シミュレーション**
    * `POST /simulations/trades` : 模擬トレード（売買）を実行
    * `GET /simulations/diaries` : トレード日記（AI自動生成）の一覧を取得

---

## 5. 画面コンポーネント構成案 (Component Architecture)

ワイヤーフレームの代替として、主要画面を構成するReactコンポーネントの階層構造と役割を定義する。

### 5.1. 共通コンポーネント (components/ui)

* `Button`: ボタン (shadcn/uiベース)
* `Card`: 情報表示用のカード (shadcn/uiベース)
* `Dialog`: モーダルダイアログ (shadcn/uiベース)
* `Spinner`: ローディングインジケータ
* `Layout`: ヘッダー、サイドバー（またはフッター）、コンテンツエリアを含む基本レイアウト

### 5.2. 主要画面・コンポーネント

#### `App` (ルート)
* 認証状態の管理 (AuthContext)
* ルーティング (React Router)
    * `/login` -> `LoginPage`
    * `/register` -> `RegisterPage`
    * `/app` (認証後) -> `DashboardLayout`

#### `DashboardLayout` (認証後のメインレイアウト)
* `Header`: 共通ヘッダー（ロゴ、ユーザーアイコン、通知）
* `Sidebar`: ナビゲーションメニュー（ダッシュボード, 学習, 模擬トレード, ニュース, リスク）
* `ContentArea`: メインコンテンツ（各ページのコンポーネントが描画される）

---

#### A. `DashboardPage` ( /app/dashboard )
* ユーザーのメインページ。
* **`WelcomeHeader`**: 「こんにちは、○○さん」
* **`DailyLessonCard`**: (✅ 4) 今日の3分レッスンへの導線。
    * `LessonProgress`: カリキュラム全体の進捗バー。
* **`PersonalizedNewsFeed`**: (✅ 3) AI要約ニュースのヘッドライン（3〜5件）。
    * `NewsCard` (クリックで詳細モーダル)
* **`MyStatusCard`**: (✅ 1) ユーザーのAI分析結果（リスク許容度、学習タイプ）の概要。

#### B. `LearningPage` ( /app/learn )
* インタラクティブ学習（チャート、クイズ）のメインページ。
* **`ChartCoachView`**: (✅ 2) AIコーチ付きチャート学習機能。
    * `InteractiveChart`: (TradingViewラッパー) チャート表示エリア。クリックイベントをハンドルする。
    * `AICoachChatBox`: AIコーチの解説が表示されるチャット風UI。
        * ユーザーがチャートをクリックすると、`InteractiveChart` がイベントを発火 -> `LearningPage` がAPIに問い合わせ -> 結果を `AICoachChatBox` に渡す。
* **`CurriculumTracker`**: (✅ 1) パーソナルカリキュラムの全容と進捗を表示するリスト。

#### C. `SimulationPage` ( /app/simulation )
* 模擬トレードとトレード日記のページ。
* **`SimulationTradeView`**: (✅ 5) 模擬トレード実行UI。
    * `SimulationChart`: 模擬トレード用のチャート（売買ボタン付き）。
    * `TradePanel`: 銘柄選択、数量入力、売買実行ボタン。
    * `PositionList`: 現在保有中の仮想ポジション一覧。
* **`TradeDiaryFeed`**: (✅ 6) AI自動生成トレード日記。
    * `DiaryCard`: AIによる分析（良かった点、改善点）を表示するカード。
        * `UserMemoInput`: ユーザーが追記メモを入力する欄。

#### D. `NewsPage` ( /app/news )
* パーソナライズドニュースの詳細ページ。
* **`NewsFilter`**: 難易度（小学生向け〜）、カテゴリでの絞り込み。
* **`FullNewsList`**: (✅ 3) AI要約ニュースの完全なリスト。
    * `NewsArticleView`: 選択されたニュースの詳細。
        * `SummarizedContent`: AI要約本文。
        * `AIInsightBox`: 「市場への影響予測」などのAI洞察。
        * `OriginalLink`: 元記事へのリンク。

#### E. `RiskDashboardPage` ( /app/risk )
* 暗号資産リスクダッシュボード。
* **`CryptoRiskViewer`**: (✅ 7)
    * `VolatilityChart`: ボラティリティの時系列グラフ。
    * `RiskReturnScatter`: リスク・リターン（シャープレシオ）の散布図。
* `CrashSimulator`: 「もし暴落が起きたら？」のシミュレーション結果。
    * `AIPersonalAlert`: 「あなたのリスク許容度に対し、この銘柄は...」というAIアラート。

---

## 6. 競合分析と本アプリのポジショニング

### 6.1. 既存の投資学習・シミュレーションアプリの分類

既存の競合サービスは、主に以下の4つのカテゴリに分類できる。

#### A. 教材提供型（eラーニング型）
* **特徴:** 体系化された投資知識（用語集、テクニカル分析、ファンダメンタルズ分析）を動画や記事で提供する。
* **例:** Udemy, Courseraの投資コース, Schoo, 証券会社が提供する学習コンテンツ
* **弱点:**
    * **画一的:** 全員に同じ内容を提供するため、ユーザーのレベルや理解度に合わないと挫折しやすい。
    * **一方通行:** インプット中心で、アウトプット（実践）の場が少ない。

#### B. 模擬トレード型（シミュレーション型）
* **特徴:** 仮想資金を使って、実際の市場（または過去の市場）で売買を体験できる。
* **例:** 多くのFX/証券会社が提供するデモトレードツール, StockSim, Investopedia Simulator
* **弱点:**
    * **「なぜ」が抜ける:** ツール（道具）の提供に留まり、「なぜその判断が良かった/悪かったのか」のフィードバックが手薄。
    * **ゲーム化しすぎる:** 仮想資金ゆえに非現実的なトレード（全額一点買いなど）になりがちで、実践的な学習にならないケースがある。

#### C. ニュース・情報提供型
* **特徴:** 経済ニュース、決算情報、アナリストレポートなどを集約して提供する。
* **例:** Bloomberg, Reuters, Investing.com, 各種ニュースアプリ
* **弱点:**
    * **情報過多:** 初心者にはノイズが多く、どの情報が重要か判断できない。
    * **難解:** 専門用語が多く、理解のハードルが高い。

#### D. トレード日記・分析型
* **特徴:** 実際のトレード履歴（または手入力）を記録し、損益や勝率をグラフ化・分析する。
* **例:** 各種トレード記録アプリ
* **弱点:**
    * **入力が面倒:** 多くの初心者は日記をつける習慣がなく、継続しない。
    * **中〜上級者向け:** すでにトレードを始めている人向けで、学習段階のユーザーには適さない。

---

### 6.2. 本アプリの強みと差別化（ブルー・オーシャン）

本アプリ「AI投資学習コンパニオン」は、上記A〜Dの全ての領域にAIによる「パーソナライズ」と「フィードバック」を持ち込むことで、既存サービスの弱点を克服し、独自のポジションを確立する。

#### 🎯 圧倒的な差別化要因

1.  **「全員違う」学習体験 (VS 教材提供型)**
    * 競合が「同じ教材」を出すのに対し、本アプリはAIが「リスク許容度」や「学習スタイル」を分析し、**個人専用カリキュラム**を生成する (✅ 1, 4)。
    * これにより、「難しすぎてついていけない」という最大の挫折要因を排除する。

2.  **「目の前のチャート」を解析 (VS 教材提供型)**
    * 競合が「チャートの“見方”」を文章で教えるのに対し、本アプリはユーザーが「今、目の前でタップしたチャート」をAIがリアルタイムで解析し、**文脈に応じたコーチング**を行う (✅ 2)。
    * これは「教科書」と「専属家庭教師」の違いに等しい。

3.  **「自分向け」ニュース (VS ニュース提供型)**
    * 競合が「同じニュース」を全員に配信するのに対し、本アプリはAIがニュースを「ユーザーの理解度」に合わせて要約し、「市場への影響予測」まで付与する (✅ 3)。
    * 情報収集の質と効率が劇的に向上する。

4.  **「フィードバック付き」シミュレーション (VS 模擬トレード型)**
    * 競合が「ツール」を提供するだけに対し、本アプリは模擬トレード後にAIが「なぜその判断に至ったか」「あなたの癖は何か」を分析し、**フィードバック**を提供する (✅ 5)。

5.  **「自動生成」トレード日記 (VS トレード日記型)**
    * 競合が「手入力」を求めるのに対し、本アプリはAIが「トレード履歴」から自動で日記（反省点、改善点）を生成する (✅ 6)。
    * 「振り返り」という最も重要だが最も面倒な作業を自動化することで、ユーザーの成長を加速させる。

### 6.3. 結論：目指すべきポジション

本アプリが目指すのは、単なる「教材」や「ツール」ではなく、
**“ユーザーにとって投資学習の伴走者になるAI”**
である。

| 方針 | 効果 (競合優位性) |
| :--- | :--- |
| ① AIが"自分専用"教材を作る | 学習負荷が激減 (競合Aの弱点を克服) |
| ② 実際のチャートやニュースで学ぶ | 実践感が楽しい (競合A, Cの弱点を克服) |
| ③ AIコーチが"個別フィードバック" | 伸びを実感できる (競合Bの弱点を克服) |
| ④ 1日3分の習慣レッスン | 挫折しにくい (継続率の向上) |
| ⑤ トレード日記自動生成 | 成長の可視化 (競合Dの弱点を克服) |

---

## 7. セットアップと実行手順

### 前提条件

以下のソフトウェアがインストールされている必要があります：

- Docker & Docker Compose
- Git
- （オプション）Go 1.21+、Node.js 18+（ローカル開発の場合）

### 7.1. Dockerを使用した起動（推奨）

#### 手順1: リポジトリをクローン

```bash
git clone https://github.com/GenkiNakashima/moneyst.git
cd moneyst
```

#### 手順2: 環境変数ファイルを作成

```bash
# ルートディレクトリに.envファイルを作成（オプション）
cp .env.example .env

# バックエンドディレクトリに.envファイルを作成
cp backend/.env.example backend/.env
```

**注意:** `.env`ファイル内の`JWT_SECRET`は本番環境では必ず変更してください。

#### 手順3: Dockerコンテナを起動

```bash
# すべてのサービス（PostgreSQL, Backend, Frontend）を起動
docker-compose up -d

# ログを確認
docker-compose logs -f
```

**開発用設定について:**
- `docker-compose.yml`は開発用に最適化されています
- バックエンド: ソースコードの変更を反映するため、volumeマウント + `go run`を使用
- フロントエンド: ホットリロードが有効で、コード変更が自動反映されます
- 初回起動時やコード変更後は、コンテナの再ビルドが必要な場合があります:
  ```bash
  docker-compose build
  docker-compose up -d
  ```

#### 手順4: アプリケーションにアクセス

- **フロントエンド:** http://localhost:3000
- **バックエンドAPI:** http://localhost:8080
- **データベース:** localhost:5432

#### サービスの停止

```bash
# サービスを停止
docker-compose down

# データベースのボリュームも削除する場合
docker-compose down -v
```

### 7.2. ローカル開発環境での起動

#### A. データベース（PostgreSQL）の起動

```bash
# Dockerでデータベースのみ起動
docker-compose up -d postgres

# または、ローカルのPostgreSQLを使用する場合は、データベースを作成
createdb moneyst
```

#### B. バックエンド（Go）の起動

```bash
cd backend

# 依存関係をインストール
go mod download

# 環境変数を設定（.envファイルを作成）
cp .env.example .env

# サーバーを起動
go run main.go
```

バックエンドは http://localhost:8080 で起動します。

#### C. フロントエンド（React + Vite）の起動

```bash
cd frontend

# 依存関係をインストール
npm install

# 開発サーバーを起動
npm run dev
```

フロントエンドは http://localhost:3000 で起動します。

### 7.3. 初回セットアップ（ユーザー登録）

1. ブラウザで http://localhost:3000 を開く
2. 「新規登録」ページでアカウントを作成
3. ログイン後、ダッシュボードが表示されます

### 7.4. データベースマイグレーション（手動実行の場合）

Dockerを使用している場合、マイグレーションは自動的に実行されます。
手動で実行する場合：

```bash
# PostgreSQLに接続
psql -h localhost -U postgres -d moneyst

# マイグレーションファイルを実行
\i backend/migrations/001_create_initial_schema.sql
```

または、GORMのAutoMigrate機能により、アプリケーション起動時に自動的にテーブルが作成されます。

### 7.5. API仕様（主要エンドポイント）

#### 認証

- `POST /api/auth/register` - ユーザー登録
- `POST /api/auth/login` - ログイン（JWTトークン取得）
- `GET /api/auth/me` - 認証済みユーザー情報取得

#### ユーザー

- `GET /api/users/me/profile` - プロファイル取得
- `PUT /api/users/me/profile` - プロファイル更新

#### カリキュラム

- `GET /api/curriculums/me` - パーソナルカリキュラム取得
- `POST /api/curriculums/:id/complete` - カリキュラム完了

#### デイリーレッスン

- `GET /api/lessons/daily` - 今日のレッスン取得
- `POST /api/lessons/:id/complete` - レッスン完了

#### シミュレーショントレード

- `POST /api/simulations/trades` - トレード実行
- `GET /api/simulations/trades` - トレード履歴取得
- `POST /api/simulations/trades/:id/close` - ポジションクローズ

#### トレード日記

- `GET /api/simulations/diaries` - トレード日記一覧取得
- `POST /api/simulations/trades/:id/diary` - トレード日記生成
- `PUT /api/simulations/diaries/:id` - トレード日記のメモ更新

#### ニュース

- `GET /api/news/personalized` - パーソナライズドニュース取得
- `POST /api/news/:id/read` - ニュース既読マーク

#### コミュニティ

- `GET /api/community/posts` - 投稿一覧取得
- `POST /api/community/posts` - 新規投稿作成（自動ファクトチェック実行）
- `GET /api/community/posts/:id` - 投稿詳細取得（返信を含む）
- `POST /api/community/posts/:id/replies` - 投稿に返信を追加
- `POST /api/community/posts/:id/like` - いいねをトグル（追加/削除）
- `GET /api/community/posts/trending` - 急上昇TOP3投稿取得
- `GET /api/community/posts/search?q=検索語` - 投稿を検索

### 7.6. トラブルシューティング

#### ポートがすでに使用されている

```bash
# 使用中のポートを確認
sudo lsof -i :3000
sudo lsof -i :8080
sudo lsof -i :5432

# プロセスを終了するか、docker-compose.ymlのポート番号を変更
```

#### データベース接続エラー

- PostgreSQLコンテナが起動しているか確認: `docker-compose ps`
- データベース接続情報（ホスト、ポート、ユーザー名、パスワード）が正しいか確認
- バックエンドの`.env`ファイルを確認

#### フロントエンドが起動しない

**症状**: `Vite requires Node.js version 20.19+ or 22.12+`のようなエラーが出る

**原因**: Node.jsのバージョンが古い

**対処法**:
1. **Docker環境の場合**:
   - `frontend/Dockerfile`でNode.js 20以上を使用していることを確認
   - 現在は`node:20-alpine`を使用しています
   - コンテナを再ビルド:
     ```bash
     docker-compose build frontend
     docker-compose up -d frontend
     ```

2. **ローカル環境の場合**:
   ```bash
   # node_modulesを削除して再インストール
   cd frontend
   rm -rf node_modules package-lock.json
   npm install
   npm run dev
   ```

#### バックエンドが起動しない

**症状**: `exec: "go": executable file not found in $PATH`

**原因**: docker-compose.ymlで開発用の設定（`target: builder`）が指定されていない

**対処法**:
- `docker-compose.yml`のbackendセクションに`target: builder`が含まれていることを確認
- コンテナを再ビルド:
  ```bash
  docker-compose build backend
  docker-compose up -d backend
  ```

#### Dockerビルドエラーが発生する

**症状**: `go build` が失敗する、または依存関係のダウンロードエラーが発生する

**対処法**:

1. **Docker Desktopでビルドログを確認する**
   - Docker Desktopアプリを開く
   - 左サイドバーの「Builds」タブをクリック
   - 失敗したビルドを選択すると詳細なログが表示されます
   - エラーメッセージの全文を確認できます

2. **コマンドラインで詳細なログを表示**
   ```bash
   # ビルドキャッシュをクリアして再ビルド
   docker-compose build --no-cache --progress=plain backend

   # または個別にビルド
   cd backend
   docker build --no-cache --progress=plain -t moneyst-backend .
   ```

3. **Goの依存関係の問題**

   **重要**: 現在のDockerfileは、ビルド時に自動的に`go.sum`を生成します。

   ```bash
   # ローカルでgo.sumを生成する場合（オプション）
   cd backend
   go mod tidy
   go mod verify
   git add go.sum
   git commit -m "Add go.sum file"

   # その後、再度Dockerビルド
   docker-compose build backend
   ```

   **注意**: Dockerfileは`go mod tidy`を実行するため、`go.sum`をリポジトリに
   含めなくてもビルドは成功しますが、ベストプラクティスとしては`go.sum`を
   リポジトリに含めることが推奨されます。これにより、ビルドの再現性と
   セキュリティが向上します。

4. **ネットワークの問題**
   - プロキシやファイアウォールの設定を確認
   - Docker Desktopの設定で「Resources」→「Network」を確認
   - 必要に応じてDNS設定を変更（例: 8.8.8.8）

5. **マルチステージビルドの確認**
   - 現在のDockerfileはマルチステージビルドを使用しています
   - ビルドステージ（builder）とランタイムステージが分離されています
   - これにより最終イメージのサイズが小さくなり、セキュリティも向上します

### 7.7. 本番デプロイ

本番環境にデプロイする場合は、以下の点に注意してください：

1. **環境変数の設定**
   - `JWT_SECRET`を強固なランダム文字列に変更
   - データベース認証情報を変更
   - CORS設定を本番ドメインに限定

2. **フロントエンドのビルド**
   ```bash
   cd frontend
   npm run build
   # dist/ディレクトリにビルド成果物が生成される
   ```

3. **バックエンドのビルド**
   ```bash
   cd backend
   go build -o main .
   ```

4. **HTTPS化**
   - Nginx や Caddy などのリバースプロキシを使用
   - Let's Encrypt で SSL証明書を取得

5. **データベースのバックアップ**
   - 定期的にPostgreSQLのバックアップを取得

### 7.8. 今後の実装予定

以下の機能は現在プレースホルダーであり、今後実装予定です：

- **AI統合**: OpenAI API / Gemini APIを使用したパーソナライズ機能
- **チャート機能**: TradingView Lightweight Chartsの統合
- **ニュースAPI**: 外部ニュースAPIからのリアルタイムニュース取得
- **リスク分析**: 暗号資産のボラティリティとリスク指標の可視化
- **通知機能**: WebSocketを使用したリアルタイム通知

---

## 8. ライセンス

This project is licensed under the MIT License.

## 9. お問い合わせ

質問や提案がある場合は、GitHubのIssuesまでお願いします。