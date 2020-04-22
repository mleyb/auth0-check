
package main

import (
	//"gopkg.in/auth0.v4"
	"gopkg.in/auth0.v4/management"
	"log"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"fmt"
	"os"
)

func main() {

	domain := "bluezero.eu.auth0.com"
	id := "Xlj91hQHbWvJASF8NA0AcJxRwNhpGaPP"
	secret := "_H72pQ9L0d1bqmiiAKyQuqK7QUEwrhzzBBMqSkMkkXeueioY4fh66wvzHy9BeoLw"

	m, err := management.New(domain, id, secret)
	if err != nil {
		// handle err



	}

	conns, err := m.Connection.List()
	if err != nil {
		// handle err



	}

	log.Println(conns.Total)

	for _, c := range conns.Connections {
		log.Println(c.GetName())
	}

	sess := session.Must(session.NewSession())

	svc := sns.New(sess)

	sub := "tit"
	msg := "dod"
	arn := "bot"

	result, err := svc.Publish(&sns.PublishInput{
		Subject: &sub,
        Message:  &msg,
        TopicArn: &arn,
    })
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    fmt.Println(*result.MessageId)
}