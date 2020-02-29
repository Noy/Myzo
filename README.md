## Myzo - The Monzo client for you!

#### Please note, you'll need to authenticate with monzo.
#### Please see their API for instructions:

## To use:

```go
package main

import (
	"github.com/Noy/Monzo"
    "fmt"
)

func main() {
    authMonzo := monzo.AuthMonzo{ClientID: "Your ClientID", 
                                UserID: "Your UserID", 
                                AccessToken: "YourAccessToken", 
                                AccountID: "Your AccountID", 
                                Debug:true}
	fmt.Println(authMonzo.GetTransaction("id", "merchant").Merchant)
}

// Or..

func main() {
	authMonzo := monzo.AuthMonzo{ClientID: "Your ClientID", 
                                UserID: "Your UserID", 
                                AccessToken: "YourAccessToken", 
                                AccountID: "Your AccountID", 
                                Debug:true}
    for _, transaction := range authMonzo.GetAllTransactions(5, "merchant") {
        fmt.Println(transaction.Merchant)
    }
}
``` 

##### More coming soon, but please note it can only be used to make..
#### applications dedicated to your account.