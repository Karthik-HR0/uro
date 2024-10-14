package main

import (
    "flag"
    "fmt"
    "net/url"
    "os"
    "strings"
    "bufio"
    "log"
)

// Main function
func main() {
    // Display logo
    displayLogo()

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

    // Call URL processing function
    processURLs(*inputFile, *outputFile)
}

// Function to display logo
func displayLogo() {
    logo := `
██╗   ██╗██████╗  ██████╗ 
██║   ██║██╔══██╗██╔═████╗
██║   ██║██████╔╝██║██╔██║
██║   ██║██╔══██╗████╔╝██║
╚██████╔╝██║  ██║╚██████╔╝
 ╚═════╝ ╚═╝  ╚═╝ ╚═════╝
URL CLEANER TOOL
`
    fmt.Fprint(os.Stdout, logo)  // Output the logo to the console
    os.Stdout.Sync()  // Ensure the logo is flushed to the output
}

// Function to process URLs from input file and write cleaned URLs to output file
func processURLs(input, output string) {
    inFile, err := os.Open(input)
    if err != nil {
        log.Fatalf("Error opening input file: %v", err)
    }
    defer inFile.Close()

    outFile, err := os.Create(output)
    if err != nil {
        log.Fatalf("Error creating output file: %v", err)
    }
    defer outFile.Close()

    scanner := bufio.NewScanner(inFile)
    writer := bufio.NewWriter(outFile)
    defer writer.Flush()

    for scanner.Scan() {
        rawURL := scanner.Text()
        cleanedURL, err := cleanURL(rawURL)
        if err != nil {
            log.Printf("Skipping invalid URL: %s\n", rawURL)
            continue
        }
        writer.WriteString(cleanedURL + "\n")
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error reading input file: %v", err)
    }

    fmt.Printf("URLs processed from %s and saved to %s\n", input, output)
}

// Function to clean and normalize a URL
func cleanURL(rawURL string) (string, error) {
    u, err := url.Parse(rawURL)
    if err != nil {
        return "", err
    }

    // Normalize scheme and host to lowercase
    u.Scheme = strings.ToLower(u.Scheme)
    u.Host = strings.ToLower(u.Host)

    // Remove fragments (if any)
    u.Fragment = ""

    // Reconstruct the URL
    return u.String(), nil
}