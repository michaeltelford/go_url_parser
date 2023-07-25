package main

import (
	"fmt"
	"strings"
)

/**
* This function counts how many unique normalized valid URLs were passed to the function
*
* Accepts a list of URLs
*
* Example:
*
* input: ['https://example.com']
* output: 1
*
* Notes:
*  - assume none of the URLs have authentication information (username, password).
*
* Normalized URL:
*  - process in which a URL is modified and standardized: https://en.wikipedia.org/wiki/URL_normalization
*
#    For example.
#    These 2 urls are the same:
#    input: ["https://example.com", "https://example.com/"]
#    output: 1
#
#    These 2 are not the same:
#    input: ["https://example.com", "http://example.com"]
#    output 2
#
#    These 2 are the same:
#    input: ["https://example.com?", "https://example.com"]
#    output: 1
#
#    These 2 are the same:
#    input: ["https://example.com?a=1&b=2", "https://example.com?b=2&a=1"]
#    output: 1
*/

func CountUniqueUrls(urls []string) int {
	uniqueUrlMap := make(map[string]bool, 0)

	for _, url := range urls {
		scheme, domain := getSchemeAndDomain(url)
		schemeAndDomain := fmt.Sprintf("%s://%s", scheme, domain)
		if _, present := uniqueUrlMap[schemeAndDomain]; !present {
			uniqueUrlMap[schemeAndDomain] = true
		}
	}

	return len(uniqueUrlMap)
}

func getSchemeAndDomain(url string) (string, string) {
	segments := strings.Split(url, "://")
	scheme := segments[0] // TODO: check out of bounds

	schemeRemoved := strings.TrimPrefix(url, fmt.Sprintf("%s://", scheme))
	domain := schemeRemoved

	// TODO: Handle if the url contains all of these characters
	if strings.Contains(domain, "/") {
		domain = removeSuffix(domain, "/")
	}
	if strings.Contains(domain, "#") {
		domain = removeSuffix(domain, "#")
	}
	if strings.Contains(domain, "?") {
		domain = removeSuffix(domain, "?")
	}

	return scheme, domain
}

func removeSuffix(url, suffix string) string {
	if strings.Contains(url, suffix) {
		index := strings.Index(url, suffix)
		suffixRemoved := url[0:index]

		return suffixRemoved
	}

	return url
}

/**
 * This function counts how many unique normalized valid URLs were passed to the function per top level domain
 *
 * A top level domain is a domain in the form of example.com. Assume all top level domains end in .com
 * subdomain.example.com is not a top level domain.
 *
 * Accepts a list of URLs
 *
 * Example:
 *
 * input: ["https://example.com"]
 * output: Hash["example.com" => 1]
 *
 * input: ["https://example.com", "https://subdomain.example.com"]
 * output: Hash["example.com" => 2]
 *
 */

func CountUniqueUrlsPerTopLevelDomain(urls []string) map[string]int {
	uniqueUrlMap := map[string]int{}

	for _, url := range urls {
		_, domain := getSchemeAndDomain(url)
		suffixRemoved := removeSuffix(domain, ".com")

		segments := strings.Split(suffixRemoved, ".")
		topLevelDomain := segments[len(segments)-1] + ".com"

		if _, present := uniqueUrlMap[topLevelDomain]; !present {
			uniqueUrlMap[topLevelDomain] = 0
		}
		uniqueUrlMap[topLevelDomain] += 1
	}

	return uniqueUrlMap
}

func main() {
	urls := []string{"https://example.com", "https://example.com/"}
	fmt.Println(CountUniqueUrls(urls))
	//    output: 1

	urls = []string{"https://example.com", "http://example.com"}
	fmt.Println(CountUniqueUrls(urls))
	//    output 2

	urls = []string{"https://example.com?", "https://example.com"}
	fmt.Println(CountUniqueUrls(urls))
	//    output: 1

	urls = []string{"https://example.com?a=1&b=2", "https://example.com?b=2&a=1"}
	fmt.Println(CountUniqueUrls(urls))
	//    output: 1

	urls = []string{"https://example.com"}
	urlMap := CountUniqueUrlsPerTopLevelDomain(urls)
	fmt.Println(urlMap)
	// output: Hash["example.com" => 1]

	urls = []string{"https://example.com", "https://subdomain.example.com"}
	urlMap = CountUniqueUrlsPerTopLevelDomain(urls)
	fmt.Println(urlMap)
	// output: Hash["example.com" => 2]
}
