package main

import (
    "flag"
    "fmt"
    "os"
    "github.com/Karthik-HR0/uro/internal/processor"
)

func main() {
    inputFile := flag.String("i", "", "File containing URLs")
    outputFile := flag.String("o", "", "Output file")
    flag.Parse()

    // Handle input file or stdin
    inputStream := os.Stdin
    if *inputFile != "" {
        file, err := os.Open(*inputFile)
        if err != nil {
            fmt.Fprintf(os.Stderr, "[ERROR] Unable to open input file: %v\n", err)
            return
        }
        defer file.Close()
        inputStream = file
    }

    // Process URLs using internal processor package
    if err := processor.ProcessURLs(inputStream, *outputFile); err != nil {
        fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
        return
    }
}
