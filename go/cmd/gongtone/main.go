package main

import (
	"flag"
	"log"
	"strconv"

	gongtone_models "github.com/fullstack-lang/gongtone/go/models"
	gongtone_stack "github.com/fullstack-lang/gongtone/go/stack"
	gongtone_static "github.com/fullstack-lang/gongtone/go/static"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	// notes := gongtone_models.GeneratePianoNotes()
	// for _, n := range notes {
	// 	log.Println(n)
	// }

	log.SetPrefix("gongtone: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := gongtone_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	stack := gongtone_stack.NewStack(r, gongtone_models.Gongtone.ToString(),
		*unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
