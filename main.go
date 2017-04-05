package main

import (
	"fmt"
	core "github.com/gerty-monit/core"
	"log"
)

func tcpMonitors() []core.Monitor {
	google := core.NewTcpMonitor("Google IP A", "Try to connect to Google port 80", "181.15.96.152", 80)
	google2 := core.NewTcpMonitor("Google IP B", "Try to connect to Google port 80", "172.217.29.3", 80)
	return []core.Monitor{google, google2}
}

func main() {
	port := 8080
	server := core.GertyServer{}
	facebook := core.NewHttpMonitor("Facebook Home", "Try to reach facebook home via Http", "http://www.facebook.com")
	server.Groups = []core.Group{
		core.Group{Name: "Network", Monitors: tcpMonitors()},
		core.Group{Name: "Http", Monitors: []core.Monitor{facebook}},
	}

	yourSlackHook := "https://hooks.slack.com/services/FOO/BAR/BAZ"
	slackAlarm := core.NewSlackAlarm(yourSlackHook)

	fromAddress := "no-reply@example.com"
	toAddress := "alarms@example.com"
	awsSmtpUser := "AWS_EMAIL_USER"
	awsSmtpPass := "AWS_EMAIL_PASS"
	dashboardHome := "https://alarms.example.com"
	emailAlarm := core.NewEmailAlarm(
		"email-smtp.us-east-1.amazonaws.com",
		"587",
		awsSmtpUser,
		awsSmtpPass,
		fromAddress,
		toAddress,
		dashboardHome,
	)

	log.Printf("alarms %s and %s disabled, uncomment below to activate",
		slackAlarm.Name(),
		emailAlarm.Name())

	// server.Alarms = []core.Alarm{
	// 	slackAlarm,
	// 	emailAlarm,
	// }

	log.Printf("server started on port %d", port)
	server.ListenAndServe(fmt.Sprintf(":%d", port))
}
