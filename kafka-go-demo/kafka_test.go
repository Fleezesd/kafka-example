package main

import "testing"

func TestConsumeMessage(t *testing.T) {
	err := ConsumeMessage()
	t.Log(err)
}

func TestLeaderConnect(t *testing.T) {
	LeaderConnect()
}

func TestTopicCreate(t *testing.T) {
	TopicCreate()
}

func TestTopicList(t *testing.T) {
	TopicList()
}
