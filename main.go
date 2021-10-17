package main

import (
	"flag"
	"fmt"
	"github.com/gregdel/pushover"
	"github.com/imroc/req"
	"time"
)

func isToday(date time.Time) bool {
	now := time.Now()
	nown := now.Unix() - now.Unix()%86400
	daten := date.Unix() - date.Unix()%86400
	return daten >= nown && daten <= nown+86400
}

func main() {
	var (
		GitHubUser       string
		GitHubApiKey     string
		PushAppKey       string
		PushRecipientKey string
		ExcludePrivate   bool
		ForceFail        bool
	)

	flag.StringVar(&GitHubUser, "u", "", "(required) GitHub User")
	flag.StringVar(&GitHubApiKey, "g", "", "(optional) GitHub API Token for private activity")
	flag.StringVar(&PushAppKey, "a", "", "(required) Pushover App Key")
	flag.StringVar(&PushRecipientKey, "r", "", "(required) Pushover Recipient Key")
	flag.BoolVar(&ExcludePrivate, "P", false, "(optional) exclude private activity")
	flag.BoolVar(&ForceFail, "F", false, "(optional) force fail to test push service")
	flag.Parse()

	if GitHubUser == "" {
		panic("GitHub User cannot be empty")
	}
	if PushAppKey == "" {
		panic("Push App Key cannot be empty")
	}
	if PushRecipientKey == "" {
		panic("Push Recipient cannot be empty")
	}

	// include api key?
	// an api key is optional, but provides access to private activities
	headers := req.Header{}
	if GitHubApiKey != "" {
		headers["Authorization"] = "token " + GitHubApiKey
	}
	fmt.Println("INF | Checking activity for", GitHubUser)

	// ForceFail: always fail to test push service
	if !ForceFail {
		// make GitHub api request
		res, err := req.Get(fmt.Sprintf("https://api.github.com/users/%s/events", GitHubUser), headers)
		if err != nil {
			panic(err)
		}

		// parse the response in the most weird way you can imagine...
		var activities []*TypedEvent
		if err = res.ToJSON(&activities); err != nil {
			panic(err)
		}

		// check activities
		for _, a := range activities {
			if !a.Public && ExcludePrivate {
				continue
			}
			// check if activity is from today
			if a.IsProductive() && isToday(a.CreatedAt) && !ForceFail {
				fmt.Println("INF | Found activity from today! :)")
				return
			}
		}
	} else {
		fmt.Println("INF | Force failing because specified by flag")
	}
	fmt.Println("WARN :: Found no activity for today!")

	app := pushover.New(PushAppKey)
	recipient := pushover.NewRecipient(PushRecipientKey)
	message := &pushover.Message{
		Message:  "Hey! 👋 You did nothing productive today!",
		Title:    "365 Days of Code 🎉",
		URL:      "https://github.com/" + GitHubUser,
		URLTitle: "GitHub/" + GitHubUser,
	}
	var (
		resp *pushover.Response
		err  error
	)
	if resp, err = app.SendMessage(message, recipient); err != nil {
		panic(err)
	}
	fmt.Println("INF | Message sent:")
	fmt.Println("---")
	fmt.Println(resp.String())
	fmt.Println("---")
}
