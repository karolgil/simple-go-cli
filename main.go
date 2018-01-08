package main

import ( // run first
	"fmt"
	"os"
	"strings"
	flag "github.com/ogier/pflag"
	gh "github.com/karolgil/simple-go-cli/gitHub"
)

var user string // run second
				// string variable declaration, available for same package - not python-like global

func init(){ // run third
			 // treated as app/package "constructor"
	flag.StringVarP(&user, "user", "u", "", "Specify username")
}

func main(){ // run fourth,
			 // special - entrypoint for app
	flag.Parse()

	if userNotProvided() {
		printUsageHelp()
		os.Exit(1)
	}

	users := getUsersListFromFlag()

	for _, u := range users {
		user := gh.GetUser(u)
		fmt.Printf("Login: %s\n", user.Login)
		fmt.Printf("ID: %s\n", user.ID)
		fmt.Printf("Email: %s\n", user.Email)
		fmt.Printf("Company: %s\n", user.Company)
		fmt.Printf("URL: %s\n", user.URL)
		fmt.Println()
	}
}

func getUsersListFromFlag() []string {
	users := strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)
	return users
}

func printUsageHelp() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func userNotProvided() bool {
	return flag.NFlag() == 0
}
