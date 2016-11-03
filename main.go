package main

import (
	"fmt"
	"math/rand"
	"time"
)

const max = 4

type (
	blog struct {
		name string
	}
	file       int
	blogReader interface {
		collect(chan string)
	}
)

func getBlogs() []blogReader {
	return []blogReader{
		blog{name: "medium.com"},
		blog{name: "xkcd.com"},
		blog{name: "golang.org"},
		file(0),
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan string)
	blogs := getBlogs()

	for _, b := range blogs {
		go b.collect(ch)
	}

	for _ = range blogs {
		result := <-ch
		fmt.Println(result)
	}
	fmt.Println("Done")
}

func (b blog) collect(ch chan string) {
	t0 := rand.Intn(200)
	time.Sleep(time.Duration(t0) * time.Millisecond)
	ch <- fmt.Sprintf("from %s", b.name)
}

func (f file) collect(ch chan string) {
	t0 := rand.Intn(200)
	time.Sleep(time.Duration(t0) * time.Millisecond)
	ch <- fmt.Sprintf("%d", f)
}
