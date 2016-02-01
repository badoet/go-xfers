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

####### Standard Go Usage
```go
import (
  "github.com/badoet/go-xfers"
)

```


Running Tests
---
There's a test suite included.  To run it, simply run:

    go test xfers_test.go

You'll have to have set the following environment variables to run the tests:

    export XFERS_TEST_KEY=XXX (Your Xfers Sandbox key, get from: https://sandbox.xfers.io/api_tokens)
    export XFERS_NOTIFY_URL=XXX (Your server callback endpoint for Xfers server to call for transaction status)

Tests currently run in sandbox.
