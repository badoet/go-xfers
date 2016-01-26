go-xfers
========

[![Circle CI](https://circleci.com/gh/badoet/go-xfers/tree/master.svg?style=svg)](https://circleci.com/gh/badoet/go-xfers/tree/master)

Xfers api interface using golang
https://www.xfers.io/

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

    export XFERS_TEST_KEY=XXX (Your XFERS Sandbox key)

Tests currently run in sandbox.
