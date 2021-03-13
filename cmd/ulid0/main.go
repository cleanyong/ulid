package main

import (
	crand "crypto/rand"
	"fmt"
	"github.com/oklog/ulid/v2"
	"math/rand"
	"time"
)

func main()  {

	entropy := ulid.Monotonic(crand.Reader, 0)
	ulid1,_:=ulid.New(ulid.Timestamp(time.Now()), entropy)
	fmt.Println(ulid1.Time(),ulid1.String())
	bin, err := ulid1.MarshalBinary()
	if err != nil {
		fmt.Println("Error !")
	}
	ulid1.UnmarshalBinary(bin)

	t := time.Unix(1000000, 0)
	entropy1 := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	fmt.Println(ulid.MustNew(ulid.Timestamp(t), entropy1))

}
