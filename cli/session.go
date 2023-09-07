package cli


func StartCliSession(){
	
}

func prepareCliSession() {
	var err error

	regularState, err = terminal.MakeRaw(0)
	if err != nil {
		log.Fatalf("Error on cli session startup: %s. Exiting.", err.Error())
	}

	log.Println("Cli-ghost session was started. Type 'help' or '/h' for more information.")
}

func handleRequest(c)