package server

import (
	"fmt"
	"log"
	"strconv"

	"github.com/QBC8-Team7/MagicCrawler/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type CommandHandler interface {
	HandleHello(update tgbotapi.Update) error
	HandleBye(update tgbotapi.Update) error
	HandleWatchlist(update tgbotapi.Update) error
}

type Handlers struct {
	Logger *logger.AppLogger
	Bot    *tgbotapi.BotAPI
	Ad     Ad
}

type Ad struct {
	PublisherAdKey string
	PublisherID    int
	Category       string
	Author         string
	Url            string
	Title          string
	Description    string
	City           string
	Neighborhood   string
	HouseType      string
	Meterage       int
	RoomsCount     int
	Year           int
	Floor          int
	TotalFloors    int
	HasWarehouse   bool
	HasElevator    bool
	Lat            string
	Lng            string
}

var lastMessageID = make(map[int64]int)

func replaceMessage(bot *tgbotapi.BotAPI, userID int64, text string, buttons *tgbotapi.InlineKeyboardMarkup) {
	if msgID, exists := lastMessageID[userID]; exists {
		fmt.Println(msgID)
		editMsg := tgbotapi.NewEditMessageText(userID, msgID, text)
		editMsg.ReplyMarkup = buttons
		_, err := bot.Send(editMsg)
		if err != nil {
			log.Println("Error editing message:", err)
			return
		}
	} else {
		msg := tgbotapi.NewMessage(userID, text)
		msg.ReplyMarkup = buttons

		sentMsg, err := bot.Send(msg)
		if err != nil {
			log.Println("Error sending message:", err)
			return
		}
		lastMessageID[userID] = sentMsg.MessageID
	}
}

func sendCategoryButtons(bot *tgbotapi.BotAPI, userID int64) {
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Rent", "rent"),
			tgbotapi.NewInlineKeyboardButtonData("Buy", "buy"),
			tgbotapi.NewInlineKeyboardButtonData("Mortgage", "mortgage"),
		),
	)
	replaceMessage(bot, userID, "Select Ad Category", &buttons)
}

func sendHouseTypeButtons(bot *tgbotapi.BotAPI, userID int64) {
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Apartment", "apartment"),
			tgbotapi.NewInlineKeyboardButtonData("Villa", "villa"),
		),
	)
	replaceMessage(bot, userID, "Select House Type", &buttons)
}

func sendWarehouseButtons(bot *tgbotapi.BotAPI, userID int64) {
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Yes", "yes"),
			tgbotapi.NewInlineKeyboardButtonData("No", "no"),
		),
	)
	replaceMessage(bot, userID, "Does the house have a warehouse?", &buttons)
}

func sendElevatorButtons(bot *tgbotapi.BotAPI, userID int64) {
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Yes", "yes"),
			tgbotapi.NewInlineKeyboardButtonData("No", "no"),
		),
	)
	replaceMessage(bot, userID, "Does the house have an elevator?", &buttons)
}

func isAdComplete(ad *Ad) bool {
	// TODO: check if ad is OK
	return true
}

func handleCallbackQuery(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	userID := update.CallbackQuery.Message.Chat.ID
	context, ok := userContext[userID]
	if !ok {
		return
	}

	ad := context.CurrentAd
	switch context.Progress {
	case 0: // Category
		ad.Category = update.CallbackQuery.Data
		sendHouseTypeButtons(bot, userID)
		context.Progress++

	case 1: // HouseType
		ad.Category = update.CallbackQuery.Data
		sendWarehouseButtons(bot, userID)
		context.Progress++

	case 2: // Warehouse
		ad.Category = update.CallbackQuery.Data
		sendElevatorButtons(bot, userID)
		context.Progress++

	case 3: // Elevator
		ad.Category = update.CallbackQuery.Data
		bot.Send(tgbotapi.NewMessage(userID, "Enter Publisher Ad Key"))
		context.Progress++
	}

	bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
}

func handleUserMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update, userID int64) {
	context, inProgress := userContext[userID]
	if !inProgress {
		bot.Send(tgbotapi.NewMessage(userID, "Use /addhouse or /updatehouse to start."))
		return
	}

	ad := context.CurrentAd
	text := update.Message.Text

	switch context.Progress {
	case 4:
		// TODO: validate publisher ad key
		if text != "" {
			ad.PublisherAdKey = text

			bot.Send(tgbotapi.NewMessage(userID, "Enter Publisher ID"))
			context.Progress++
		} else {
			bot.Send(tgbotapi.NewMessage(userID, "Invalid value. Enter Publisher Ad Key again"))
		}
	case 5: // PublisherID
		// TODO: validation
		pid, err := strconv.Atoi(text)
		if err == nil && pid > 0 {
			ad.PublisherID = pid
			bot.Send(tgbotapi.NewMessage(userID, "Enter Author"))
			context.Progress++
		} else {
			bot.Send(tgbotapi.NewMessage(userID, "Invalid value. Enter Publisher ID again"))
		}
	case 6: // Author
		// TODO: validation
		if text != "" {
			ad.Author = text
			bot.Send(tgbotapi.NewMessage(userID, "Enter Title"))
			context.Progress++
		} else {
			bot.Send(tgbotapi.NewMessage(userID, "Invalid value. Enter Author again"))
		}
	case 7: // Title
		// TODO: validation
		if text != "" {
			ad.Title = text
			bot.Send(tgbotapi.NewMessage(userID, "Enter Description"))
			context.Progress++
		} else {
			bot.Send(tgbotapi.NewMessage(userID, "Invalid value. Enter Title again"))
		}
	case 8: // Description
		// TODO: validation
		if text != "" {
			ad.Description = text
			bot.Send(tgbotapi.NewMessage(userID, "Enter City"))
			context.Progress++
		} else {
			bot.Send(tgbotapi.NewMessage(userID, "Invalid value. Enter Description again"))
		}
	case 9: // City
		// TODO: validation
		if text != "" {
			ad.City = text
			bot.Send(tgbotapi.NewMessage(userID, "Enter Neighborhood"))
			context.Progress++
		} else {
			bot.Send(tgbotapi.NewMessage(userID, "Invalid value. Enter City again"))
		}
	case 10: // Neighborhood
		// TODO: validation
		if text != "" {
			ad.Neighborhood = text
			bot.Send(tgbotapi.NewMessage(userID, "Enter Meterage"))
			context.Progress++
		} else {
			bot.Send(tgbotapi.NewMessage(userID, "Invalid value. Enter Neighborhood again"))
		}
	case 11: // Meterage
		// TODO: validation
		meterage, err := strconv.Atoi(text)
		if err == nil && meterage >= 0 {
			ad.Meterage = meterage
			bot.Send(tgbotapi.NewMessage(userID, "Enter Rooms Count"))
			context.Progress++
		} else {
			bot.Send(tgbotapi.NewMessage(userID, "Invalid value. Enter Meterage again"))
		}
	case 12: // RoomsCount
		// TODO: validation
		rooms, err := strconv.Atoi(text)
		if err == nil && rooms >= 0 {
			ad.RoomsCount = rooms
			bot.Send(tgbotapi.NewMessage(userID, "Enter Year"))
			context.Progress++
		} else {
			bot.Send(tgbotapi.NewMessage(userID, "Invalid value. Enter Rooms Count again"))
		}

	case 13: // Year
		// TODO: validation
		year, err := strconv.Atoi(text)
		if err == nil && year >= 1250 {
			ad.Year = year
			bot.Send(tgbotapi.NewMessage(userID, "Enter Floor number"))
			context.Progress++
		} else {
			bot.Send(tgbotapi.NewMessage(userID, "Invalid value. Enter year again"))
		}
	case 14: // Floor
		// TODO: validation
		floor, err := strconv.Atoi(text)
		if err == nil && floor >= 0 {
			ad.Floor = floor
			bot.Send(tgbotapi.NewMessage(userID, "Enter Total Floors"))
			context.Progress++
		} else {
			bot.Send(tgbotapi.NewMessage(userID, "Invalid value. Enter Floor number again"))
		}
	case 15: // TotalFloors
		// TODO: validation
		floors, err := strconv.Atoi(text)
		if err == nil && floors >= 0 {
			ad.TotalFloors = floors
			bot.Send(tgbotapi.NewMessage(userID, "Does it have a warehouse?"))
			context.Progress++
		} else {
			bot.Send(tgbotapi.NewMessage(userID, "Invalid value. Enter Total Floors again"))
		}
	}
}
