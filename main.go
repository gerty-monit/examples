package main

import (
	"fmt"
	"github.com/gerty-monit/core"
	a "github.com/gerty-monit/core/alarms"
	m "github.com/gerty-monit/core/monitors"
	"log"
)

func tcpMonitors() []m.Monitor {
	google := m.NewTcpMonitor("Google IP A", "Try to connect to Google port 80", "181.15.96.152", 80)
	google2 := m.NewTcpMonitor("Google IP B", "Try to connect to Google port 80", "172.217.29.3", 80)
	return []m.Monitor{google, google2}
}

func main() {
	port := 8080
	server := gerty.GertyServer{}
	facebook := m.NewHttpMonitor("Facebook Home", "Try to reach facebook home via Http", "http://www.facebook.com")
	server.Groups = []m.Group{
		m.Group{Name: "Network", Monitors: tcpMonitors()},
		m.Group{Name: "Http", Monitors: []m.Monitor{facebook}},
	}

	yourSlackHook := "https://hooks.slack.com/services/FOO/BAR/BAZ"
	slackAlarm := a.NewSlackAlarm(yourSlackHook)

	fromAddress := "no-reply@example.com"
	toAddress := "alarms@example.com"
	awsSmtpUser := "AWS_EMAIL_USER"
	awsSmtpPass := "AWS_EMAIL_PASS"
	dashboardHome := "https://alarms.example.com"
	emailAlarm := a.NewEmailAlarm(
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

	// server.Alarms = []a.Alarm{
	// 	slackAlarm,
	// 	emailAlarm,
	// }

	log.Printf("server started on port %d", port)
	server.ListenAndServe(fmt.Sprintf(":%d", port))
}
