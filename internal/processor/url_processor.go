package processor

import (
    "bufio"
    "fmt"
    "net/url"
    "os"
    "strings"
    "github.com/Karthik-HR0/uro/pkg/utils"
)

var urlMap = make(map[string]map[string]map[string]string)

func ProcessURLs(inputStream *os.File, outputFile string) error {
    scanner := bufio.NewScanner(inputStream)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if strings.HasSuffix(line, "/") {
            line = strings.TrimSuffix(line, "/")
        }
        parsedURL, err := url.Parse(line)
        if err == nil && parsedURL.Host != "" {
            processURL(parsedURL)
        }
    }
    return outputResults(outputFile)
}

func processURL(parsedURL *url.URL) {
    host := parsedURL.Host
    path := parsedURL.Path
    params := utils.ParamsToDict(parsedURL.RawQuery)

    if _, exists := urlMap[host]; !exists {
        urlMap[host] = make(map[string]map[string]string)
    }

    if isNewPath(path, params, urlMap[host]) {
        urlMap[host][path] = params
    }
}

func outputResults(outputFile string) error {
    var output *os.File
    var err error
    if outputFile != "" {
        output, err = os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
        if err != nil {
            return fmt.Errorf("unable to open output file: %v", err)
        }
        defer output.Close()
    } else {
        output = os.Stdout
    }

    for host, paths := range urlMap {
        for path, params := range paths {
            fmt.Fprintf(output, "https://%s%s%s\n", host, path, utils.DictToParams(params))
        }
    }
    return nil
}

func isNewPath(path string, params map[string]string, pathMap map[string]map[string]string) bool {
    existingParams, exists := pathMap[path]
    if !exists {
        return true
    }
    return !utils.SameParams(existingParams, params)
}
