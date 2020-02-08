/* 
I will now use my excellent Go skills to write excellent test


*/

package main
import (
	"testing"
	"encoding/json"
)

// Some sort of brilliant test to discern any json response from API
func JSONString(str string) bool {
    var jsonStr string
    err := json.Unmarshal([]byte(str), &jsonStr)
    return err == nil
}

func isJSON(s string) bool {
    var js map[string]interface{}
    return json.Unmarshal([]byte(s), &js) == nil
}


// Test Create request returns as expected on accounts resource




// Test Fetch request returns as expected on accounts resource
func TestFetchReq(t *testing.T) {
	resp := FetchReq()
	if ( !isJSON(resp) ) {
		t.Error("Fetch does not return JSON")
	}
}


// Test List request returns as expected on accounts resource

// Test Delete request returns as expected on accounts resource