package hydradblayer

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkSQLDBReads(b *testing.B) {
	fmt.Println("test...")
	dblayer, err := ConnectDatabase("mysql", "root:root@/hydra")
	if err != nil {
		b.Fatal("Could not connect to Hydra chat system: ", err)
	}
	findMembersMB(b, dblayer)
}

func findMembersMB(b *testing.B, dblayer DBLayer) {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(dblayer.FindMember(1))
}
