package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    // Define input and output flags
    inputFile := flag.String("i", "", "Input file containing URLs")
    outputFile := flag.String("o", "", "Output file for cleaned URLs")
    
    // Define help flag
    help := flag.Bool("h", false, "Display help")

    // Parse the flags
    flag.Parse()

    // Show help message if -h is provided
    if *help {
        fmt.Println("Usage of uro:")
        fmt.Println("  -i string")
        fmt.Println("        Input file containing URLs")
        fmt.Println("  -o string")
        fmt.Println("        Output file for cleaned URLs")
        fmt.Println("  -h")
        fmt.Println("        Display help")
        os.Exit(0)
    }

    // Ensure input and output files are provided
    if *inputFile == "" || *outputFile == "" {
        fmt.Println("Error: Input and output files must be specified.")
        fmt.Println("Use -h flag for help.")
        os.Exit(1)
    }

    // Call your URL processing function here (replace with actual function)
    processURLs(*inputFile, *outputFile)
}

// Placeholder for URL processing logic
func processURLs(input, output string) {
    fmt.Printf("Processing URLs from %s and saving to %s...\n", input, output)
    // Add your URL normalization logic here
}
