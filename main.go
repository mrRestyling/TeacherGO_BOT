package main

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// const TOKEN = os.getenv('TOKEN')
const CONTACT = "https://t.me/mr_restyling"

var bot *tgbotapi.BotAPI
var chatID int64

// *Описание для telegram*
// start - Запустить помощника
// contacts - Связаться с создателем

func connectWithTelegram() {
	var err error
	if bot, err = tgbotapi.NewBotAPI(TOKEN); err != nil {
		panic("Cannot connect to Telegram")
	}

}

func StartMessage() {

	// Send a sticker in response
	sticker := tgbotapi.NewSticker(chatID, tgbotapi.FileID("CAACAgIAAxkBAAEMpodmvfcCsjZBj9iFZvd4SjaDuTXKJQAChwIAAladvQpC7XQrQFfQkDUE"))
	bot.Send(sticker)

	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Виды сортировок"),

			tgbotapi.NewKeyboardButton("Поиск факториала"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("binaryToDecimal"),
			tgbotapi.NewKeyboardButton("Дополнительно"),
		),
	)

	message := `Привет! 
Я твой личный помощник в изучении языка программирования GO.

Вот несколько примеров:`
	msgConfig := tgbotapi.NewMessage(chatID, message)
	msgConfig.ReplyMarkup = keyboard
	bot.Send(msgConfig)
}

func SendContacts() {
	contact := tgbotapi.NewMessage(chatID, CONTACT)
	bot.Send(contact)
}

func SortsGO() {

	sticker := tgbotapi.NewSticker(chatID, tgbotapi.FileID("CAACAgIAAxkBAAEMpn5mvfOEhzxgDS6_lzGXB48kulcNswACIhMAAl1scEuuLYe9O-OAPjUE"))
	bot.Send(sticker)

	time.Sleep(time.Second * 2)

	link := `
	Тут можно ознакомиться подробнее:
	
	https://github.com/mrRestyling/SortsGO
	`

	text := `Виды сортировок:
	1) Пузырьковая сортировка
	2) Сортировка выбором
	3) Сортировка вставкой
	4) Быстрая сортировка
	5) Сортировка слиянием
`
	textBot := tgbotapi.NewMessage(chatID, text)
	bot.Send(textBot)

	time.Sleep(time.Second * 1)
	linkBot := tgbotapi.NewMessage(chatID, link)
	bot.Send(linkBot)
}

func Factorial() {

	sticker := tgbotapi.NewSticker(chatID, tgbotapi.FileID("CAACAgIAAxkBAAEMpn5mvfOEhzxgDS6_lzGXB48kulcNswACIhMAAl1scEuuLYe9O-OAPjUE"))
	bot.Send(sticker)

	time.Sleep(time.Second * 2)

	funcText := `Функция факториала: 

	func factorial (num uint) uint { 
      if num == 0 {
          return 1
      }
		return num * factorial(num-1)
	}
	`

	text := funcText + "\n" + `
	===========================
	Описание функции:

	uint 
	- "unsigned integer" (беззнаковое целое число)

	if num == 0 
	- Это базовый случай рекурсии. Факториал 0 равен 1

	return num * factorial(num-1) 
	- Здесь функция вызывает саму себя с аргументом num-1.
	Это рекурсивный шаг, где функция умножает num на результат вызова factorial(num-1).
	Рекурсия продолжается до тех пор, пока не дойдет до базового случая
	` + "\n"
	textBot := tgbotapi.NewMessage(chatID, text)
	bot.Send(textBot)

	link := `
	Тут можно ознакомиться подробнее:
	
	https://github.com/mrRestyling/SortsGO
	`

	time.Sleep(time.Second * 1)
	linkBot := tgbotapi.NewMessage(chatID, link)
	bot.Send(linkBot)
}

func mem() {
	sticker := tgbotapi.NewSticker(chatID, tgbotapi.FileID("CAACAgIAAxkBAAEMpn5mvfOEhzxgDS6_lzGXB48kulcNswACIhMAAl1scEuuLYe9O-OAPjUE"))
	bot.Send(sticker)
}

func binaryToDecimal() {
	sticker := tgbotapi.NewSticker(chatID, tgbotapi.FileID("CAACAgIAAxkBAAEMpn5mvfOEhzxgDS6_lzGXB48kulcNswACIhMAAl1scEuuLYe9O-OAPjUE"))
	bot.Send(sticker)

	time.Sleep(time.Second * 2)

	funcText := `Функция перевода двоичного числа в десятичное: 

	func binToDec2() {
          var bin string = "1010"
          n, err := strconv.ParseInt(bin, 2, 64)
          if err != nil {
              fmt.Println("ошибка парсинга")
          return
	}
	fmt.Println(n)
}
	`

	text := funcText + "\n" + `
	===========================
	Описание функции:

	strconv.ParseInt(bin, 2, 64)

	- ParseInt функция для перевода двоичного числа в десятичное из пакета strconv

	- Принимает двоичное число в виде строки(bin), основание системы счисления и размер бита для результата

	` + "\n"
	textBot := tgbotapi.NewMessage(chatID, text)
	bot.Send(textBot)

	link := `
	Остальные примеры можно посмотреть здесь:
	
	https://github.com/mrRestyling/ExamplesGo/blob/main/binaryToDecimal/main.go
	`

	time.Sleep(time.Second * 1)
	linkBot := tgbotapi.NewMessage(chatID, link)
	bot.Send(linkBot)
}

func main() {

	connectWithTelegram()

	updateConfig := tgbotapi.NewUpdate(0) // 0 - сколько сообщений пропустить чтобы добраться до нужного

	// получаем специальный канал связи с помощью updateConfig + проходимся циклом
	// for - особенность обхода канала (не for _, _ := range)
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message != nil && update.Message.Text == "/start" {
			chatID = update.Message.Chat.ID
			StartMessage()
		} else if update.Message != nil && update.Message.Text == "/contacts" {
			chatID = update.Message.Chat.ID
			SendContacts()
		} else if update.Message != nil && update.Message.Text == "Виды сортировок" {
			chatID = update.Message.Chat.ID
			SortsGO()
		} else if update.Message != nil && update.Message.Text == "Поиск факториала" {
			chatID = update.Message.Chat.ID
			Factorial()
		} else if update.Message != nil && update.Message.Text == "binaryToDecimal" {
			chatID = update.Message.Chat.ID
			binaryToDecimal()
		} else if update.Message != nil && update.Message.Text == "Дополнительно" {
			chatID = update.Message.Chat.ID
			mem()
		}
	}

}
