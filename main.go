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
	// make GitHub api request
	fmt.Println("INF | Checking activity for", GitHubUser)
	res, err := req.Get(fmt.Sprintf("https://api.github.com/users/%s/events", GitHubUser), headers)
	if err != nil {
		panic(err)
	}

	// parse the response in the most weird way you can imagine...
	var raw []interface{}
	if err = res.ToJSON(&raw); err != nil {
		panic(err)
	}
	var activities []CodeActivity
	if activities, err = ParseActivities(raw); err != nil {
		panic(err)
	}

	// check activities
	for _, a := range activities {
		// check if activity is from today
		if isToday(a.GetCreatedAt()) && !ForceFail {
			return
		}
		if !a.IsPublic() && ExcludePrivate {
			continue
		}
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
	var resp *pushover.Response
	if resp, err = app.SendMessage(message, recipient); err != nil {
		panic(err)
	}
	fmt.Printf("message sent! %+v\n", resp)
}
