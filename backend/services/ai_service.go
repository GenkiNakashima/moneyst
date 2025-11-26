package services

import (
	"fmt"
	"strings"
)

// AIService handles AI-related operations
type AIService struct {
	// In production, this would connect to OpenAI, Gemini, or other AI APIs
	// For now, we'll implement placeholder logic
}

// NewAIService creates a new AI service instance
func NewAIService() *AIService {
	return &AIService{}
}

// FactCheckContent performs AI-based fact-checking on content
// Returns: (factCheckResult string, status string, error)
// Status: "approved", "flagged", "removed"
func (s *AIService) FactCheckContent(content string) (string, string, error) {
	// TODO: Integrate with actual AI API (OpenAI, Gemini, etc.)
	// For now, implement basic keyword-based checks

	content = strings.ToLower(content)

	// Check for obvious misinformation keywords
	flaggedKeywords := []string{
		"100% guaranteed profit",
		"risk-free investment",
		"get rich quick",
		"this stock will definitely",
		"insider information",
	}

	for _, keyword := range flaggedKeywords {
		if strings.Contains(content, keyword) {
			result := fmt.Sprintf("⚠️ この投稿には誤解を招く可能性のある表現が含まれています: '%s'。投資には常にリスクが伴います。", keyword)
			return result, "flagged", nil
		}
	}

	// Check for content that should be removed
	bannedKeywords := []string{
		"scam",
		"ponzi",
		"pump and dump",
	}

	for _, keyword := range bannedKeywords {
		if strings.Contains(content, keyword) {
			result := fmt.Sprintf("この投稿は不適切なコンテンツ('%s')が含まれているため、削除されました。", keyword)
			return result, "removed", nil
		}
	}

	// If no issues found
	return "✓ 投稿内容を確認しました。明らかな誤情報は検出されませんでした。", "approved", nil
}

// GenerateAIResponse generates an AI response to a user's question
// This is triggered when a post or reply contains @checkAI
func (s *AIService) GenerateAIResponse(question string) (string, error) {
	// TODO: Integrate with actual AI API (OpenAI, Gemini, etc.)
	// For now, provide a template response

	question = strings.TrimSpace(strings.ReplaceAll(question, "@checkAI", ""))

	if question == "" {
		return "こんにちは！@checkAIです。投資に関する質問があれば、お気軽にお尋ねください。", nil
	}

	// Basic response template
	response := fmt.Sprintf(`🤖 **AI Assistant (@checkAI)**

ご質問: %s

【回答】
申し訳ございませんが、現在AI APIの統合は開発中です。以下のガイドラインを参考にしてください：

1. **投資の基本原則**: 分散投資とリスク管理が重要です
2. **情報の確認**: 複数の信頼できる情報源から確認しましょう
3. **専門家への相談**: 重要な投資判断は専門家に相談することをお勧めします

より詳しい情報が必要な場合は、学習ページやシミュレーション機能をご利用ください。

*注: これはプレースホルダーの応答です。本番環境ではOpenAI/Gemini APIと統合されます。*`, question)

	return response, nil
}

// CheckIfAIInvoked checks if content contains @checkAI mention
func (s *AIService) CheckIfAIInvoked(content string) bool {
	return strings.Contains(strings.ToLower(content), "@checkai")
}
