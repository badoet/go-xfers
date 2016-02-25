go-xfers
========

[![Circle CI](https://circleci.com/gh/badoet/go-xfers/tree/master.svg?style=svg)](https://circleci.com/gh/badoet/go-xfers/tree/master)

Xfers api interface using golang
https://www.xfers.io/

Xfers api documentation:
http://xfers.github.io/docs/

There are more api in the documentations that has yet to be implemented.
This library only use the minimum required api to perform transaction with Xfers

Xfers Version : 3.0

Feature:
- Create Charge
- Verify Charge
- List Charges
- Get Account Info
- Comes with predefined structs to store Xfers server responses


go get github.com/badoet/go-xfers

####### Standard Go Usage (EXAMPLE)
```go
import (
   "github.com/badoet/go-xfers"
)

var XfersAPI *xfers.Client

func init() {
	if IsProduction() {
		XfersIsSandbox = false // if is production, force to non-sandbox
	} else {
		XfersIsSandbox = true
	}

	var err error

	if XfersIsSandbox {
		XfersAPI, err = xfers.NewClient(XFERS_KEY_SANDBOX, true)
	} else {
		XfersAPI, err = xfers.NewClient(XFERS_KEY, false)
	}

	if err != nil {
		panic(err.Error())
	}
}
```

In this example, we created a file to contain the XfersClient
This client will survive throughout the app lifetime and can be reused to perform multiple actions
In normal circumstances, there is no need to recreate this client for every single action

####### Create an Xfers Charge (EXAMPLE)
```go
func XfersCheckout(order Order) (xfers.Charge, error) {

	xfersItemList := []xfers.Item{}
	for _, cartItem := range cartList { // for reference only
		xfersItem := xfers.Item{}
		xfersItem.Price = cartItem.Price
		xfersItem.Quantity = 1
		xfersItem.Name = cartItem.Name
		xfersItemList = append(xfersItemList, xfersItem)
	}

	chargeParam := xfers.ChargeReqParam{}
	chargeParam.Amount = order.TotalCost
	chargeParam.Shipping = order.ShippingFee
	chargeParam.Currency = order.Currency
	chargeParam.OrderID = order.Id
	chargeParam.NotifyURL = <URL for Xfers server to call to notify the transaction status>
	chargeParam.ReturnURL = <URL for user to go back to after successful transaction>
	chargeParam.CancelURL = <URL for user to go back to if user cancel the transaction>
	chargeParam.Redirect = false // false means the Xfers server will return us JSON response instead of redirecting directly to the Xfers page
	chargeParam.Refundable = true
	chargeParam.ReceiptEmail = order.Email
	chargeParam.HrsToExpirations = 48
	chargeParam.Items = xfersItemList

	xfersCharge, err := XfersAPI.CreateCharge(chargeParam)

	return xfersCharge, err
}
```

This is the basic example on how to create the Xfers Charge
The xfersCharge will contain `Url` property that you need to send your user to.
They will make the payment based on the instructions stated in the specified Xfers page.

After user make a payment, Xfers will send a POST request to the `NotifyUrl` stated in the above example
To capture the request, you can do something like this:

####### Accepting Xfers Notification (EXAMPLE)
```go
func XfersDone(c *gin.Context) { // im using Gin for my webframework in this example
	defer c.String(http.StatusOK, "ok")
	notification := xfers.NotifyParam{}
	err := c.Bind(&notification)

	// verify the xfers request (reccomended that you do this process)
	isVerified, err := XfersVerifyNotification(notification)

	// update order status
	if notification.Status == "paid" {
		// payment successful! update the order
	} else {

	}

```

In the `notification.Status`, it can be one of these values: “cancelled” or “paid” or “expired”.
The value of the status should be pretty self explainatory.


As for the `XfersVerifyNotification` function, you can follow this code snippet:

####### Verify Xfers Notification (EXAMPLE)
```go
func XfersVerifyNotification(params xfers.NotifyParam) (bool, error) {
	verifyParam := xfers.VerifyParam{}
	verifyParam.Init(params)
	xfersMsg, err := XfersAPI.VerifyCharge(params.TxnId, verifyParam)
	return (xfersMsg.Msg == "VERIFIED"), err
}

```

Feel free to open New Issue, if you encounter error or need help with the integration
Looking forward for feedbacks. Hope someone will find this useful!
Currently used in production at https://www.ikoustyle.com

Running Tests
---
There's a test suite included.  To run it, simply run:

    go test xfers_test.go

You'll have to have set the following environment variables to run the tests:

    export XFERS_TEST_KEY=XXX (Your Xfers Sandbox key, get from: https://sandbox.xfers.io/api_tokens)
    export XFERS_NOTIFY_URL=XXX (Your server callback endpoint for Xfers server to call for transaction status)

Tests currently run in sandbox.
