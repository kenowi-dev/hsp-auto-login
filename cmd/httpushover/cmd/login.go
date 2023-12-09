package cmd

import (
	"github.com/kenowi-dev/hspscraper"
	"github.com/spf13/cobra"
	"log"
	"time"
)

type LoginConfig struct {
	sport        string
	courseNumber string
	email        string
	password     string
	retry        int
	date         string
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to HSP-Course",
	Long:  "Automatically login to a given Hochschulsport Course",
	Args:  cobra.NoArgs,
	Run:   runLogin,
}
var loginConfig LoginConfig

func init() {

	flags := loginCmd.Flags()

	flags.StringVarP(&loginConfig.sport, "sport", "s", "", "The name of the sport. Can also be the url to the sport.")
	flags.StringVarP(&loginConfig.courseNumber, "courseNumber", "c", "", "The course number")
	flags.StringVarP(&loginConfig.date, "date", "d", "", "The date of the course")
	flags.StringVarP(&loginConfig.email, "email", "e", "", "The email for the login")
	flags.StringVarP(&loginConfig.password, "password", "p", "", "The password for the login")
	flags.IntVarP(&loginConfig.retry, "retry", "r", -1, "Interval between looking if the course is open. Will not retry if omitted.")

	_ = loginCmd.MarkFlagRequired("sport")
	_ = loginCmd.MarkFlagRequired("courseNumber")
	_ = loginCmd.MarkFlagRequired("email")
	_ = loginCmd.MarkFlagRequired("password")
	_ = loginCmd.MarkFlagRequired("date")
}

func runLogin(_ *cobra.Command, _ []string) {

	date, err := time.Parse(time.DateOnly, loginConfig.date)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = hspscraper.Register(loginConfig.sport, loginConfig.courseNumber, loginConfig.email, loginConfig.password, date)
	if err == nil {
		log.Println("Registered")
		return
	}

	log.Println(err.Error())
	if loginConfig.retry > 0 {
		err = retryLogin(loginConfig.sport, loginConfig.courseNumber, loginConfig.email, loginConfig.password, date, loginConfig.retry)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}

func retryLogin(sport, courseNumber, email, password string, date time.Time, retry int) error {
	ticker := time.NewTicker(time.Duration(retry) * time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		err := hspscraper.Register(sport, courseNumber, email, password, date)
		if err == nil {
			return nil
		}
	}
}
