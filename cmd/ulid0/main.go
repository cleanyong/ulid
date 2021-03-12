package main

import (
	crand "crypto/rand"
	"fmt"
	"github.com/oklog/ulid/v2"
	"time"
)

func main()  {

	entropy := ulid.Monotonic(crand.Reader, 0)
	ulid,_:=ulid.New(ulid.Timestamp(time.Now()), entropy)
	fmt.Println(ulid.Time(),ulid.String())
}
