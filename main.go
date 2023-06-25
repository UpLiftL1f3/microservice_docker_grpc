package main

import (
	"flag"
)

func main() {
	// client := client.New("http://localhost:3000")

	// price, err := client.FetchPrice(context.Background(), "ET")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// fmt.Printf("%+v\n", price)

	// return
	listenAddr := flag.String("listenAddr", ":3000", "the port the service is running on")
	flag.Parse()

	svc := NewLoggingService(NewMetricsService(&priceService{}))

	server := NewJSONAPIServer(*listenAddr, svc)

	server.Run()

	// fmt.Println(price)
}
