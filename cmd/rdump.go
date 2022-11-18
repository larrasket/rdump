package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"sort"

	"github.com/salehmu/rdump/internal/py"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Please pass a username as command line argument")
		return
	}
	u := os.Args[1]
	cmd := exec.Command("python", "-c", py.Script, "-u", u)
	_, _ = cmd.Output()
	if _, err := os.Stat(fmt.Sprintf("%s.shelf", u)); errors.Is(err, os.ErrNotExist) {
		log.Fatal(`Couldn't find the database file, please make sure that:
					1. You have the right permissions to write for the current working directory
	 				2. You entered a proper reddit username`)
	}
	des := fmt.Sprintf(py.Descriptor, u)
	cmd = exec.Command("python", "-c", des)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	data := string(output)
	mp := map[string]int{}
	type kv struct {
		Key   string
		Value int
	}
	r := regexp.MustCompile(`\br/(..*?\b)`)
	subreddits := r.FindAllString(data, -1)
	for _, value := range subreddits {
		mp[value]++
	}

	var ss []kv
	for k, v := range mp {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	fmt.Printf("All subreddit %s is involved in:(sorted by frequency)\n\n", u)
	for _, value := range ss {
		fmt.Printf("https://www.reddit.com/%s\n", value.Key)
	}

}
