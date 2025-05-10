package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	// Example 1: Parse a URL
	rawURL := "https://example.com:8080/path?name=john&age=25#section1"
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// Accessing different parts of the URL
	fmt.Println("Scheme:", parsedURL.Scheme)     // https
	fmt.Println("Host:", parsedURL.Host)         // example.com:8080
	fmt.Println("Path:", parsedURL.Path)         // /path
	fmt.Println("RawQuery:", parsedURL.RawQuery) // name=john&age=25
	fmt.Println("Fragment:", parsedURL.Fragment) // section1

	// Example 2: Working with Query Parameters
	queryParams := parsedURL.Query()
	fmt.Println("Name:", queryParams.Get("name")) // john
	fmt.Println("Age:", queryParams.Get("age"))   // 25

	// Example 3: Building a URL
	newURL := &url.URL{
		Scheme:   "https",
		Host:     "api.example.com",
		Path:     "/users",
		RawQuery: "active=true",
	}
	fmt.Println("Built URL:", newURL.String())

	// Example 4: URL Encoding/Decoding
	encodedString := url.QueryEscape("Hello World! @#$%")
	fmt.Println("Encoded:", encodedString)

	decodedString, err := url.QueryUnescape(encodedString)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println("Decoded:", decodedString)

	// Example 5: Joining URL paths
	base := "https://api.example.com"
	endpoint := "/users/profile"
	joinedURL := strings.TrimRight(base, "/") + "/" + strings.TrimLeft(endpoint, "/")
	fmt.Println("Joined URL:", joinedURL)

	// Example 6: Working with Query Parameters (Advanced)
	params := url.Values{}
	params.Add("name", "john")
	params.Add("age", "25")
	params.Add("hobbies", "reading")
	params.Add("hobbies", "gaming") // Multiple values for same key

	fmt.Println("All hobbies:", params["hobbies"]) // [reading gaming]
	fmt.Println("Query string:", params.Encode())  // age=25&hobbies=reading&hobbies=gaming&name=john

	// Example 7: Resolving relative URLs
	base2, _ := url.Parse("https://example.com/base/")
	relative, _ := url.Parse("../images/photo.jpg")
	resolved := base2.ResolveReference(relative)
	fmt.Println("Resolved URL:", resolved.String())
}
