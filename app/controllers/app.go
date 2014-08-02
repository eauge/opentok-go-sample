package controllers

import (
	"fmt"
	"github.com/eauge/opentok"
	"github.com/revel/revel"
	"os"
	"strconv"
)

var (
	apiKey     = readIntVariable("API_KEY", true)
	apiSecret  = readStringVariable("API_SECRET", true)
	ot         = opentok.OpenTok{ApiKey: apiKey, ApiSecret: apiSecret}
	options    = opentok.SessionOptions{}
	properties = opentok.TokenProperties{}
	session    *opentok.Session
)

type App struct {
	*revel.Controller
}

type Message struct {
	ApiKey    int
	SessionId string
	Token     string
}

func (c App) Index() revel.Result {

	session = getSession()

	// The session is created when the first
	// participant joins
	token, err := session.GenerateToken(properties)

	if err != nil {
		panic("Token could not be created")
	}

	m := Message{
		SessionId: session.Id,
		Token:     string(token),
		ApiKey:    apiKey,
	}

	return c.Render(m)
}

func getSession() *opentok.Session {
	if session == nil {
		var err error
		session, err = opentok.NewSession(ot, options)
		if err != nil {
			panic("Session could not be created")
		}

		if err = session.Create(); err != nil {
			panic("Session could not be created")
		}
	}
	return session
}

func readIntVariable(variable string, mandatory bool) int {
	value := os.Getenv(variable)
	if len(value) == 0 && mandatory {
		panic(fmt.Sprintf("Environment variable : %s expected", variable))
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("Environment variable : %s was expected to be an integer : %s", variable, err))
	}
	return intValue
}

func readStringVariable(variable string, mandatory bool) string {
	value := os.Getenv(variable)
	if len(value) == 0 && mandatory {
		panic(fmt.Sprintf("Environment variable : %s expected", variable))
	}

	return value
}
