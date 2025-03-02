package validator

import (
	"fmt"
	"net/http"
)

// Validate external links
func ValidateExternalLink(link string) {
	fmt.Println("🌐 Validating external link:", link)
	resp, err := http.Head(link)
	if err != nil {
		fmt.Println("❌ Broken external link:", link, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		fmt.Println("❌ Invalid external link:", link, "Status:", resp.StatusCode)
	} else {
		// fmt.Println("Valid external link:", link)
	}
}
