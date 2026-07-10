package telegram

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// SendToAdmin Yangi kelgan she'rni Telegram bot orqali sizga yuboradi
func SendToAdmin(author, text string) error {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	if botToken == "" || chatID == "" {
		return fmt.Errorf("telegram sozlamalari topilmadi")
	}

	// Telegram xabari formati
	message := fmt.Sprintf("📝 *Yangi Kelin Salom keldi!*\n\n*Kimdan:* %s\n\n*Matn:* \n%s", author, text)
	
	// Telegram API URL
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	// So'rov yuborish
	resp, err := http.PostForm(apiURL, url.Values{
		"chat_id":    {chatID},
		"text":       {message},
		"parse_mode": {"Markdown"},
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram xato qaytardi: %d", resp.StatusCode)
	}

	return nil
}