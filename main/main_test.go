/* 
I will now use my excellent Go skills to write excellent test


*/

package main
import (
	"testing"
	"encoding/json"
)

// Some sort of brilliant test to discern any response from API
func JSONString(str string) bool {
    var jsonStr string
    err := json.Unmarshal([]byte(str), &jsonStr)
    return err == nil
}

// Test Create request returns as expected on accounts resource




// Test Fetch request returns as expected on accounts resource
func TestfetchReq(t *testing.T) {
	resp := fetchReq()
	if ( !JSONString(resp) ) {
		t.Error("Fetch does not return JSON")
	}
}


// Test List request returns as expected on accounts resource

// Test Delete request returns as expected on accounts resource