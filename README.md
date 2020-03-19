## Myzo - The Monzo client for you!

#### Please note, you'll need to authenticate with monzo.
#### Please see their API for instructions:

## To use:

```go
package main

import (
    "fmt"
    "github.com/Noy/Myzo"
)

func main() {
    auth := myzo.Myzo{ClientID: "Your ClientID", 
                                UserID: "Your UserID", 
                                AccessToken: "YourAccessToken", 
                                AccountIDs: map[string]string{"Personal":"ID", "Business":"ID", "Joint":"ID"}, 
                                Debug:true}
    currentAccountID := auth.AccountIDs["Personal"]
	fmt.Println(auth.GetTransaction("id", "merchant", currentAccountID).Merchant)
}

// Or..

func main() {
	auth := myzo.Myzo{ClientID: "Your ClientID", 
                                UserID: "Your UserID", 
                                AccessToken: "YourAccessToken", 
                                AccountIDs: map[string]string{"Personal":"ID", "Business":"ID", "Joint":"ID"}, 
                                Debug:true}
    currentAccountID := auth.AccountIDs["Personal"]
    for _, transaction := range auth.GetAllTransactions(5, 0, "", currentAccountID) {
        fmt.Println(transaction.Merchant)
    }
}
``` 

##### More coming soon, but please note it can only be used to make applications dedicated to your account.