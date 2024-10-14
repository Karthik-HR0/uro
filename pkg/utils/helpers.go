package utils

import (
    "fmt"
    "sort"
    "strings"
)

func ParamsToDict(query string) map[string]string {
    params := make(map[string]string)
    for _, param := range strings.Split(query, "&") {
        kv := strings.SplitN(param, "=", 2)
        if len(kv) == 2 {
            params[kv[0]] = kv[1]
        } else {
            params[kv[0]] = ""
        }
    }
    return params
}

func DictToParams(params map[string]string) string {
    var parts []string
    for key, value := range params {
        parts = append(parts, fmt.Sprintf("%s=%s", key, value))
    }
    sort.Strings(parts)
    return "?" + strings.Join(parts, "&")
}

func SameParams(params1, params2 map[string]string) bool {
    if len(params1) != len(params2) {
        return false
    }
    for key, val1 := range params1 {
        if val2, ok := params2[key]; !ok || val1 != val2 {
            return false
        }
    }
    return true
}
