



# Im Going around completing online tutorials to show that I have been learning many Golang Concepts

## I have many incomplete project but from here on I will start projects with the intention of completing them moving forwards! this will improve my github presence and show to show that I really do put the time in to better myself and the code I write.

### i wanted to keep the filing format the same as AnthonyGG on youtube to better follow along but i am aware that it is a bit messy

### It took me a second to understand what was happening in this video but I think it VERY interesting.

- Anthony made a service in service.go and made a interface with a method called FetchPrice and  empty struct called priceFetcher
- when he added the method FetchPrice to the empty struct he implicitly attacked the PriceFetcher interface to the priceFetcher struct because they both express a method of FetchPrice

- NEXT 
- he created the logging.go file which contains a logger
- the logger has a struct with a field called next which contains a PriceFetcher interface as the type expressing the next field
- this means that next has a method attacked to it called FetchPrice which can be called by the new loggingService struct
- whenever you call FetchPrice using the loggingService it runs through the logger witch will display the err, the returned price and the time it took to run the function! (Very cool)

- AFTER 
- anthony created a metrics service with the struct metricsService which contains a field called next that has a value of the PriceFetcher interface that expresses the FetchPrice method
- he creates a FetchPriceMethod which he adds to the metricsService 
- he uses this FetchPrice function as a "forward" tool that allows him to access the FetchPrice function from the service then log metrics towards a third party and then finally send the response to the logger which will get the time, the price, and any errors

- LATER
- he creates the api.go file
- ge creates the JSONAPIServer struct which holds two fields the port (listenAddr) and the service of type PriceFetcher. The JSONAPIServer holds two methods which is handleFetchPrice() and Run().
- the Run() is what runs the when the api receives a request to the "/" route which then adds the requestID to the context and runs the apiFn that was passed in which is handleFetchPrice()
- the method handleFetchPrice() reaches into the passed in struct which is PriceFetcher and runs the method FetchPrice() and return the values as a PriceResponse{}