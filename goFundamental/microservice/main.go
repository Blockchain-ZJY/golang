package main

import "flag"

func main() {

	// client mode:

	// client := client.New("http://localhost:3000")
	// price, err := client.FetchPrice(context.Background(), "ET1H")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v\n", price)
	// return

	// server mode:

	listenAddr := flag.String("listenAddr", ":3000", "the listen address of the server")
	flag.Parse()
	svc := NewLoggingService(&priceFetcher{})
	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()
}
