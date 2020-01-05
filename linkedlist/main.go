package main

import (
	gears "github.com/algo/linkedlist/engine"
	"time"
)

func main() {
	rightNow := time.Now().Unix()
	f := &gears.Feed{}
	p1 := gears.Post{
		Body:        "Post 1",
		PublishDate: rightNow,
	}
	p2 := gears.Post{
		Body:        "Post 2",
		PublishDate: rightNow + 10,
	}
	p3 := gears.Post{
		Body:        "Post 3",
		PublishDate: rightNow + 20,
	}
	p4 := gears.Post{
		Body:        "Post 4",
		PublishDate: rightNow + 30,
	}
	f.Append(&p1)
	f.Append(&p2)
	f.Append(&p3)
	f.Append(&p4)


	newPost := &gears.Post{
		Body:        "This is a new post",
		PublishDate: rightNow + 15,
	}
	f.Insert(newPost)
	f.Inspect()
}
