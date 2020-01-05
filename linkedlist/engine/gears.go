package engine

import (
	"errors"
	"fmt"
)

type Post struct {
	Body         string
	PublishDate int64 // Unix timestamp
	Next *Post // link to the Next Post
}

type Feed struct {
	Length int
	Start  *Post
	End    *Post
}

func (f *Feed) AppEndSlow(newPost *Post) {
	if f.Length == 0 {
		f.Start = newPost
	} else {
		currentPost := f.Start
		for currentPost.Next != nil {
			currentPost = currentPost.Next
		}
		currentPost.Next = newPost
	}
	f.Length++
}

func (f *Feed) Append(newPost *Post) {
	if f.Length == 0 {
		f.Start = newPost
		f.End = newPost
	} else {
		lastPost := f.End
		lastPost.Next = newPost
		f.End = newPost
	}
	f.Length++
}

func (f *Feed) Remove(PublishDate int64) {
	if f.Length == 0 {
		panic(errors.New("Feed is empty"))
	}

	var previousPost *Post
	currentPost := f.Start

	for currentPost.PublishDate != PublishDate {
		if currentPost.Next == nil {
			panic(errors.New("No such Post found."))
		}

		previousPost = currentPost
		currentPost = currentPost.Next
	}
	previousPost.Next = currentPost.Next

	f.Length--
}

func (f *Feed) Insert(newPost *Post) {
	if f.Length == 0 {
		f.Start = newPost
	} else {
		var previousPost *Post
		currentPost := f.Start

		for currentPost.PublishDate < newPost.PublishDate {
			previousPost = currentPost
			currentPost = previousPost.Next
		}

		previousPost.Next = newPost
		newPost.Next = currentPost
	}
	f.Length++
}

func (f *Feed) Inspect() {
	if f.Length == 0 {
		fmt.Println("Feed is empty")
	}
	fmt.Println("========================")
	fmt.Println("Feed Length: ", f.Length)

	currentIndex := 0
	currentPost := f.Start

	for currentIndex < f.Length {
		fmt.Printf("Item: %v - %v\n", currentIndex, currentPost)
		currentPost = currentPost.Next
		currentIndex++
	}
	fmt.Println("========================")
}