package uuid

import (
	"fmt"
	"testing"
)

func TestGenerateOrderedUUID(t *testing.T) {
	uuid := GenerateOrderedUUID()
	fmt.Println(uuid)
	if len(uuid) != 36 {
		t.Errorf("error generating uuid")
	}
}
