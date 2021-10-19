package main

import (
	"flag"
	"fmt"
	"github.com/gregdel/pushover"
	"github.com/imroc/req"
	"os"
	"strings"
	"time"
)

func isToday(date time.Time) bool {
	now := time.Now()
	nown := now.Unix() - now.Unix()%86400
	daten := date.Unix() - date.Unix()%86400
	return daten >= nown && daten <= nown+86400
}

const (
	DefaultMessageText  = "Hey! 👋 You did nothing productive today!"
	DefaultMessageTitle = "365 Days of Code 🎉"
)

func main() {
	var (
		GitHubUser       = flag.String("u", "", "(required) GitHub User")
		PushAppKey       = flag.String("a", "", "(required) Pushover App Key")
		PushRecipientKey = flag.String("r", "", "(required) Pushover Recipient Key")
		GitHubApiKey     = flag.String("g", "", "(optional) GitHub API Token for private activity")
		ExcludePrivate   = flag.Bool("P", false, "(optional) exclude private activity")
		ForceFail        = flag.Bool("F", false, "(optional) force fail to test push service")
		MessageText      = flag.String("mText", DefaultMessageText, "(optional) Message Text")
		MessageTitle     = flag.String("mTitle", DefaultMessageTitle, "(optional) Message Title")
	)
	flag.Parse()

	// from env
	env(GitHubUser, "GITHUB_USER")
	env(GitHubApiKey, "GITHUB_PAT")
	env(PushAppKey, "PUSHOVER_APP_KEY")
	env(PushRecipientKey, "PUSHOVER_RECIPIENT_KEY")
	env(MessageText, "MESSAGE_TEXT")
	env(MessageTitle, "MESSAGE_TITLE")

	// check required flags
	if *GitHubUser == "" {
		panic("GitHub User cannot be empty")
	}
	if *PushAppKey == "" {
		panic("Push App Key cannot be empty")
	}
	if *PushRecipientKey == "" {
		panic("Push Recipient cannot be empty")
	}

	// include api key?
	// an api key is optional, but provides access to private activities
	headers := req.Header{}
	if *GitHubApiKey != "" {
		headers["Authorization"] = "token " + *GitHubApiKey
	}
	fmt.Println("INFO | Checking activity for", *GitHubUser)

	// ForceFail: always fail to test push service
	if !*ForceFail {
		// make GitHub api request
		res, err := req.Get(fmt.Sprintf("https://api.github.com/users/%s/events", *GitHubUser), headers)
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
			if !a.Public && *ExcludePrivate {
				continue
			}
			// check if activity is from today
			if a.IsProductive() && isToday(a.CreatedAt) && !*ForceFail {
				fmt.Println("INFO | Found activity from today! :)")
				return
			}
		}
	} else {
		fmt.Println("INFO | Force failing because specified by flag")
	}
	fmt.Println("WARN | Found no activity for today!")

	app := pushover.New(*PushAppKey)
	recipient := pushover.NewRecipient(*PushRecipientKey)
	message := &pushover.Message{
		Message:  *MessageText,
		Title:    *MessageTitle,
		URL:      "https://github.com/" + *GitHubUser,
		URLTitle: "GitHub/" + *GitHubUser,
	}
	var (
		resp *pushover.Response
		err  error
	)
	if resp, err = app.SendMessage(message, recipient); err != nil {
		panic(err)
	}
	fmt.Println("INFO | Message sent:")
	wrap(resp.String())
}

func wrap(msg string) {
	m := 1
	for _, l := range strings.Split(msg, "\n") {
		if x := len(l); x > m {
			m = x
		}
	}
	fmt.Println(strings.Repeat("-", m))
	fmt.Println(msg)
	fmt.Println(strings.Repeat("-", m))
}

func env(flag *string, key string) *string {
	// flag already set?
	if *flag == "" {
		if x, ok := os.LookupEnv(key); ok {
			*flag = x
		}
	}
	return flag
}
