package main

import "flag"

type Config struct {
	CollectionPath         string
	DestinationPath        string
	ApibFileName           string
	ForceApibCreation      bool
	ForceResponsesCreation bool
	DumpRequest            string
}

func (c *Config) Init() {
	flag.StringVar(&c.CollectionPath, "collection", "", "Path to the Postman collection export")
	flag.StringVar(&c.CollectionPath, "c", "", "Path to the Postman collection export")

	flag.StringVar(&c.DestinationPath, "destination", "./", "Destination folder path for the generated files")
	flag.StringVar(&c.DestinationPath, "d", "./", "Destination folder path for the generated files")

	flag.StringVar(&c.ApibFileName, "apibname", "", "Set a custom name for the generated .apib file")

	flag.BoolVar(&c.ForceApibCreation, "force-apib", false, "Override existing .apib files")
	flag.BoolVar(&c.ForceResponsesCreation, "force-responses", false, "Override existing response files")

	flag.StringVar(&c.DumpRequest, "dump-request", "", "Output the markup for a single request. (Takes a request name)")

	flag.Parse()
}
