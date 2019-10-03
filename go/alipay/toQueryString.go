package alipay

import (
	"fmt"
	"sort"
	"strings"
)


func (m M) ToQueryString(ignoreEmpty bool, sortByASCII bool) string {
	var data []string
	for k, v := range m {
		if ignoreEmpty == true {
			if v != "" {
				data = append(data, fmt.Sprintf("%s=%s", k, v))
			}
		} else {
			data = append(data, fmt.Sprintf("%s=%s", k, v))
		}
	}
	if sortByASCII {
		sort.Strings(data)
	}
	qs := strings.Join(data, "&")
	return qs
}
