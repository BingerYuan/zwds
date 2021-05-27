package ziwei

import (
	"fmt"
	"testing"
)

func TestGetZiWei(t *testing.T) {
	wxn := 6
	lday := 1
	s := GetZiWei(wxn, lday)
	fmt.Println(s)
}
