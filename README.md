 # gRPC Word Search Challenge
 This app enables the user to search and affect a set of pre-defined search terms as below.

 ## Useful Commands
 `make help`  
 *  prints out list and description of all available make commands    

 `make proto`
 *  Generates proto api, grpc-gateway definition, and swagger definition    

 `make build`  
 *  Runs test suite, and builds binaries for osx, linux, and windows, and places them in the ./bin folder.    
 
 `make buildrun`  
 *  Runs tests, builds binaries and runs osx version of binary    
 
 `make run`  
 *  Runs osx version of binary    
 
 `make generate`  
 *  Generates proto output, runs tests, builds binaries, and runs osx version of binary    

 ### Single word search
 ```
 curl "http://localhost:8090/api/words?word=sawyer"
 ```
 
 ### Update word search list
 ```
 curl -X "POST" "http://localhost:8090/api/words/sawyer"
 ```
 
 ### Fetch top 5 most searched for words
 ```
 curl "http://localhost:8090/api/words/popular"
 ```

 ## Points to note
 I usually use the **moq** library by Mat Ryer to mock my interfaces and test them.  
 [https://github.com/matryer/moq](https://github.com/matryer/moq) however when running the *moq* command 
 on the command line, or when trying to run `go generate` to generate my mocks it would fail.
 This is believe is due to `gofmt` being run under the hood and not liking some of the
 stuff it is seeing in the auto-generated proto files. If time allowed I would use another library like
 [https://github.com/stretchr/testify](https://github.com/stretchr/testify) to test my interfaces, or even write
 my own mock implementations of them however, this week at home as been pretty distracting. I'm more than
 happy to show examples of my tests in other repo's or create a couple of interfaces in a new app and test them for you.

 Any questions please don't hesitate to get in touch.  
 justindavidson23@gmail.com