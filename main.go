package main

import (
	"fmt"
	gh "github.com/karolgil/simple-go-cli/gitHub"
	flag "github.com/ogier/pflag"
	"os"
	"strconv"
	"strings"
	"sync"
)

var usersFlag string
var maxWorkers int

func init() {
	maxProcs, err := strconv.Atoi(os.Getenv("GOMAXPROCS"))
	if err != nil {
		maxProcs = 1
	}

	flag.StringVarP(&usersFlag, "usersFlag", "u", "", "Specify username")
	flag.IntVarP(&maxWorkers, "max-workers", "m", maxProcs, "Specify maximum number of workers")
}

func main() {
	flag.Parse()

	if userNotProvided() {
		printUsageHelp()
		os.Exit(1)
	}

	usernames := make(chan string, maxWorkers)
	results := make(chan gh.User)
	var tasksWg, resultsWg sync.WaitGroup

	go getUsersListFromFlag(usernames)

	for i := 0; i < maxWorkers; i++ {
		tasksWg.Add(1)
		go gh.GetUsers(&tasksWg, usernames, results)
	}

	resultsWg.Add(1)
	go presentResults(&resultsWg, results)

	tasksWg.Wait()
	close(results)
	resultsWg.Wait()
}

func presentResults(wg *sync.WaitGroup, results <-chan gh.User) {
	defer wg.Done()
	for ghUser := range results {
		fmt.Printf("Login: %s\n", ghUser.Login)
		fmt.Printf("ID: %s\n", ghUser.ID)
		fmt.Printf("Email: %s\n", ghUser.Email)
		fmt.Printf("Company: %s\n", ghUser.Company)
		fmt.Printf("URL: %s\n", ghUser.URL)
		fmt.Println()
	}
}

func getUsersListFromFlag(usernames chan<- string) {
	for _, username := range strings.Split(usersFlag, ",") {
		usernames <- username
	}
	close(usernames)
}

func printUsageHelp() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func userNotProvided() bool {
	return flag.NFlag() == 0
}
