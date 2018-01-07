package main

import ( // run first
	"fmt"
	flag "github.com/ogier/pflag"
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
	fmt.Println(user)
}
