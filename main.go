package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	githubToken := os.Getenv("GITHUB_TOKEN")
	client, ctx := initClient(githubToken)

	notifs, resp, err := client.Activity.ListNotifications(ctx, &github.NotificationListOptions{
		All: true,
	})

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		panic(resp.Status)
	}

	for _, notif := range notifs {
		fmt.Printf("%s | %s\n", notif.Repository.GetName(), notif.GetSubject().GetTitle())
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func initClient(githubToken string) (client *github.Client, ctx context.Context) {
	ctx = context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client = github.NewClient(tc)

	return client, ctx
}
