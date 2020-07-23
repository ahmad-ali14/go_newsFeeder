package newsFeed_test

import (
	"testing"

	"newsFeeder/platform/newsFeed"
)

func TestAdd(t *testing.T) {

	feed := newsFeed.New()

	feed.Add(newsFeed.Item{"hello", "this is my first post ever"})

	feed.Add(newsFeed.Item{})

	if len(feed.Items) != 2 {
		t.Errorf("item was not added")
	}

}

func TestGetAll(t *testing.T) {

	feed := newsFeed.New()

	feed.Add(newsFeed.Item{"hello", "this is my first post ever"})

	feed.Add(newsFeed.Item{})

	result := feed.GetAll()

	if len(result) != 2 {
		t.Errorf("getAll is not working")
	}

}
