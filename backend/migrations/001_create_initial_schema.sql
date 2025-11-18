-- 001_create_initial_schema.sql
-- Initial database schema for AI Investment Learning Companion

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    hashed_password VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    risk_tolerance VARCHAR(50) CHECK (risk_tolerance IN ('low', 'medium', 'high')),
    learning_style VARCHAR(50) CHECK (learning_style IN ('text', 'visual', 'conversational')),
    investment_experience VARCHAR(50) CHECK (investment_experience IN ('none', 'beginner', 'intermediate')),
    trading_pattern_memo TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Lessons table (educational content master)
CREATE TABLE IF NOT EXISTS lessons (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    content_body TEXT NOT NULL,
    difficulty_level INTEGER CHECK (difficulty_level BETWEEN 1 AND 5),
    category VARCHAR(100) CHECK (category IN ('chart', 'glossary', 'economy', 'quiz')),
    estimated_time_min INTEGER,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Curriculums table (personalized learning progress)
CREATE TABLE IF NOT EXISTS curriculums (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    module_type VARCHAR(100) CHECK (module_type IN ('chart', 'news', 'simulation', 'quiz')),
    content_id UUID,
    status VARCHAR(50) CHECK (status IN ('pending', 'in_progress', 'completed')) DEFAULT 'pending',
    "order" INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Simulation trades table (mock trading history)
CREATE TABLE IF NOT EXISTS simulation_trades (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    symbol VARCHAR(50) NOT NULL,
    trade_type VARCHAR(10) CHECK (trade_type IN ('buy', 'sell')) NOT NULL,
    entry_price DECIMAL(18, 8) NOT NULL,
    quantity DECIMAL(18, 8) NOT NULL,
    entry_at TIMESTAMP WITH TIME ZONE NOT NULL,
    exit_price DECIMAL(18, 8),
    exit_at TIMESTAMP WITH TIME ZONE,
    profit_loss DECIMAL(18, 8),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Trade diaries table (AI-generated trade journals)
CREATE TABLE IF NOT EXISTS trade_diaries (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    trade_id UUID REFERENCES simulation_trades(id) ON DELETE CASCADE,
    generated_analysis TEXT,
    market_context TEXT,
    user_memo TEXT,
    diary_date DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Summarized news table (AI-summarized and personalized news)
CREATE TABLE IF NOT EXISTS summarized_news (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    original_url VARCHAR(1000),
    original_title VARCHAR(500),
    summarized_title VARCHAR(500),
    summarized_content TEXT,
    ai_insight TEXT,
    published_at TIMESTAMP WITH TIME ZONE NOT NULL,
    read_status BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- User achievements table (gamification)
CREATE TABLE IF NOT EXISTS user_achievements (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    achievement_type VARCHAR(100) NOT NULL,
    achievement_title VARCHAR(255) NOT NULL,
    achievement_description TEXT,
    earned_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Daily lessons table (3-minute daily lessons)
CREATE TABLE IF NOT EXISTS daily_lessons (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    lesson_date DATE NOT NULL,
    content TEXT NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    points_earned INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, lesson_date)
);

-- Indexes for performance
CREATE INDEX idx_curriculums_user_id ON curriculums(user_id);
CREATE INDEX idx_curriculums_status ON curriculums(status);
CREATE INDEX idx_simulation_trades_user_id ON simulation_trades(user_id);
CREATE INDEX idx_simulation_trades_entry_at ON simulation_trades(entry_at);
CREATE INDEX idx_trade_diaries_user_id ON trade_diaries(user_id);
CREATE INDEX idx_trade_diaries_diary_date ON trade_diaries(diary_date);
CREATE INDEX idx_summarized_news_user_id ON summarized_news(user_id);
CREATE INDEX idx_summarized_news_published_at ON summarized_news(published_at);
CREATE INDEX idx_daily_lessons_user_id_date ON daily_lessons(user_id, lesson_date);

-- Update timestamp trigger function
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Apply update timestamp triggers
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_lessons_updated_at BEFORE UPDATE ON lessons FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_curriculums_updated_at BEFORE UPDATE ON curriculums FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_simulation_trades_updated_at BEFORE UPDATE ON simulation_trades FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_trade_diaries_updated_at BEFORE UPDATE ON trade_diaries FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_summarized_news_updated_at BEFORE UPDATE ON summarized_news FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
