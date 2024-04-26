package main

// 3.3 masalani ishlashini ko'rsatish uchun githubni o'zidan change qilindi
// dev branchdan o'zgarish amamlga oshirildi

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)


var (
	RED     = "\033[31m"
	GREEN   = "\033[32m"
	YELLOW  = "\033[33m"
	BLUE    = "\033[34m"
	MAGENTA = "\033[35m"
	CYAN    = "\033[36m"
	WHITE   = "\033[97m"
	RESET   = "\033[0m"
)

func clearOS() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	// findDifference() // 3.2 - masala

	me := User{"Azizbek", []*Event{}}

	/*
		// Agar hammasini qo'lda kiritib chiqmoqchi bo'lmasangiz bularni yoqib qo'ying
		// shunda o'zi rejalar qo'shilib qoladi va siz kerakli paytdagini print qilaverasiz :)

		bugungiReja := Event{"BugungiReja", time.Now()}
		ertangiReja := Event{"ErtangiReja", time.Now().AddDate(0, 0, 1)}
		shuHaftaReja := Event{"ShuHaftaReja", time.Now().AddDate(0, 0, 7)}
		shuOyReja := Event{"ShuOyReja", time.Now().AddDate(0, 1, 0)}
		shuYilReja := Event{"ShuYilReja", time.Now().AddDate(0, 0, 365)}
		keyingiYilReja := Event{"KeyingiYilReja", time.Now().AddDate(1, 0, 0)}

		me.AddEvent(&bugungiReja)
		me.AddEvent(&ertangiReja)
		me.AddEvent(&shuHaftaReja)
		me.AddEvent(&shuOyReja)
		me.AddEvent(&shuYilReja)
		me.AddEvent(&keyingiYilReja)

	*/

	for {
		fmt.Println(CYAN, "Siz asosiy bo'limdasiz, kerakli buyruqlardan birini tanlang: \n  0 > Dasturdan chiqish \n  1 > Hodisalani ko'rish \n  2 > Hodisa qo'shish \n  3 > Hodisani o'chirish", RESET)
		var inp int
		fmt.Scan(&inp)
		switch inp {
		case 0:
			clearOS()
			return
		case 1:
			clearOS()
			fmt.Println(CYAN, "Qaysi paytdagi hodisalarni ko'rmoqchisiz? \n  1 > Bugungi \n  2 > Ertangi \n  3 > Shu haftalik \n  4 > Shu oylik \n  5 > Shu yillik \n  6 > Barchasi \n  0 > Orqaga", RESET)
			fmt.Scan(&inp)
			switch inp {
			case 0:
				clearOS()
				continue
			case 1:
				me.PrintEvents(true, false, false, false, false, false)
			case 2:
				me.PrintEvents(false, true, false, false, false, false)
			case 3:
				me.PrintEvents(false, false, true, false, false, false)
			case 4:
				me.PrintEvents(false, false, false, true, false, false)
			case 5:
				me.PrintEvents(false, false, false, false, true, false)
			case 6:
				me.PrintEvents(false, false, false, false, false, true)
			}
		case 2:
			clearOS()
		again:
			var eventName string
			var eventDate string
			fmt.Println(CYAN, "Qo'shmoqchi bo'lgan hodisaning nomini kiriting: \n 0 > Bekor qilish", RESET)
			fmt.Scan(&eventName)
			if eventName == "0" {
				clearOS()
				continue
			}
			fmt.Println(CYAN, "Qo'shmoqchi bo'lgan hodisaning sanasi va vaqtini kiriting. 17.04.2024-23:59 kabi, boshqacha format qabul qilimaydi: \n 0 > Bekor qilish ", RESET)
			fmt.Scan(&eventDate)
			if eventDate == "0" {
				clearOS()
				continue
			}
			isValid, date := isDateValid(eventDate)
			if isValid {
				newEvent := Event{eventName, date}
				me.AddEvent(&newEvent)
				fmt.Println(GREEN, "Hodisa qo'shildi", RESET)
			} else {
				fmt.Println(RED, "Xato sana kiritdingiz! Sana 17.04.2024-23:59 kabi bo'lishi kerak.", RESET)
				goto again
			}
		case 3:
			clearOS()
			fmt.Println(CYAN, "O'chirmoqchi bo'lgan hodisaning nomini kiriting: \n 0 > Bekor qilish", RESET)
			var eventToRemove string
			fmt.Scan(&eventToRemove)
			if eventToRemove == "0" {
				clearOS()
				continue
			}
			isRemoved := me.RemoveEvent(eventToRemove)
			if isRemoved {
				clearOS()
				fmt.Println(GREEN, "Hodisa muvaffaqiyatli o'chirildi!", RESET)
			} else {
				clearOS()
				fmt.Println(RED, "Hodisani o'chirib bo'lmadi: bunday nomdagi hodisa topilmadi!", RESET)
			}
		}
	}
}

type User struct {
	Name   string
	Events []*Event
}

type Event struct {
	EventName string
	Date      time.Time
}

func (u *User) AddEvent(e *Event) {
	u.Events = append(u.Events, e)
}

func (u *User) PrintEvents(today, tomarrow, thisWeek, thisMonth, thisYear, all bool) {
	clearOS()
	fmt.Println(MAGENTA, "#--------------------------------------------------------------#\n", RESET)

	if len(u.Events) == 0 {
		fmt.Println(RED, "Hech qanday hodisalar yo'q :(", RESET)
	} else {
		switch {
		case today:
			count := 0
			for _, event := range u.Events {
				if event.Date.YearDay() == time.Now().YearDay() && event.Date.Year() == time.Now().Year() {
					count++
				}
			}
			if count > 0 {
				fmt.Println(BLUE, "Bugungi hodisalar:", RESET)
				fmt.Println(BLUE, "Hodisa nomi      Sana", RESET)
				for _, event := range u.Events {
					if event.Date.YearDay() == time.Now().YearDay() && event.Date.Year() == time.Now().Year() {
						fmt.Println(YELLOW, event.EventName, ":      ", event.Date.Format("January 02 2006, 15:04"))
					}
				}
			} else {
				fmt.Println(RED, "Hech qanday hodisalar yo'q :(", RESET)
			}
		case tomarrow:
			count := 0
			for _, event := range u.Events {
				if event.Date.YearDay() == time.Now().AddDate(0, 0, 1).YearDay() && event.Date.Year() == time.Now().Year() {
					count++
				}
			}
			if count > 0 {
				fmt.Println(BLUE, "Ertangi hodisalar:", RESET)
				fmt.Println(BLUE, "Hodisa nomi      Sana", RESET)
				for _, event := range u.Events {
					if event.Date.YearDay() == time.Now().AddDate(0, 0, 1).YearDay() && event.Date.Year() == time.Now().Year() {
						fmt.Println(YELLOW, event.EventName, ":      ", event.Date.Format("January 02 2006, 15:04"))
					}
				}
			} else {
				fmt.Println(RED, "Hech qanday hodisalar yo'q :(", RESET)
			}

		case thisWeek:
			count := 0
			for _, event := range u.Events {
				if event.Date.After(time.Now().AddDate(0, 0, -int(time.Now().Weekday()))) && event.Date.Before(time.Now().AddDate(0, 0, 7-int(time.Now().Weekday()))) {
					count++
				}
			}
			if count > 0 {
				fmt.Println(BLUE, "Bu haftadagi hodisalar:", RESET)
				fmt.Println(BLUE, "Hodisa nomi      Sana", RESET)
				for _, event := range u.Events {
					if event.Date.After(time.Now().AddDate(0, 0, -int(time.Now().Weekday()))) && event.Date.Before(time.Now().AddDate(0, 0, 7-int(time.Now().Weekday()))) {
						fmt.Println(YELLOW, event.EventName, ":      ", event.Date.Format("January 02 2006, 15:04"))
					}
				}
			} else {
				fmt.Println(RED, "Hech qanday hodisalar yo'q :(", RESET)
			}
		case thisMonth:
			count := 0
			for _, event := range u.Events {
				if event.Date.Year() == time.Now().Year() && event.Date.Month() == time.Now().Month() {
					count++
				}
			}
			if count > 0 {
				fmt.Println(BLUE, "Bu oydagi hodisalar:", RESET)
				fmt.Println(BLUE, "Hodisa nomi      Sana", RESET)
				for _, event := range u.Events {
					if event.Date.Year() == time.Now().Year() && event.Date.Month() == time.Now().Month() {
						fmt.Println(YELLOW, event.EventName, ":      ", event.Date.Format("January 02 2006, 15:04"))
					}
				}
			} else {
				fmt.Println(RED, "Hech qanday hodisalar yo'q :(", RESET)
			}
		case thisYear:
			count := 0
			for _, event := range u.Events {
				if event.Date.Year() == time.Now().Year() {
					count++
				}
			}
			if count > 0 {
				fmt.Println(BLUE, "Bu yildagi hodisalar:", RESET)
				fmt.Println(BLUE, "Hodisa nomi      Sana", RESET)
				for _, event := range u.Events {
					if event.Date.Year() == time.Now().Year() {
						fmt.Println(YELLOW, event.EventName, ":      ", event.Date.Format("January 02 2006, 15:04"))
					}
				}
			} else {
				fmt.Println(RED, "Hech qanday hodisalar yo'q :(", RESET)
			}
		case all:
			fmt.Println(BLUE, "Barcha hodisalar:", RESET)
			fmt.Println(BLUE, "Hodisa nomi      Sana", RESET)
			for _, event := range u.Events {
				fmt.Println(YELLOW, event.EventName, ":      ", event.Date.Format("January 02 2006, 15:04"))
			}
		}
	}

	fmt.Println(MAGENTA, "\n #--------------------------------------------------------------#", RESET)
}

func (u *User) RemoveEvent(eventName string) bool {
	for i, event := range u.Events {
		if event.EventName == eventName {
			u.Events = append(u.Events[:i], u.Events[i+1:]...)
			return true
		}
	}
	return false
}

func isDateValid(dateString string) (valid bool, date time.Time) {
	layout := "02.01.2006-15:04"
	date, err := time.Parse(layout, dateString)
	if err != nil {
		return false, date
	}
	return true, date
}

func findDifference() {
	now := time.Now()
	mybirth := time.Date(2024, time.November, 29, 0, 0, 0, 0, time.UTC)

	diff := mybirth.Sub(now)
	days := int(diff.Hours() / 24)
	hours := int(diff.Hours()) % 24
	minutes := int(diff.Minutes()) % 60
	seconds := int(diff.Seconds()) % 60

	fmt.Printf("%d days, %d hours, %d minutes, %d seconds left until your birthday", days, hours, minutes, seconds)
	fmt.Println()
}
