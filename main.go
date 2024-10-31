package main

import (
    "bufio"
    "flag"
    "fmt"
    "net/url"
    "os"
    "strings"
    "sync"
    "time"

    "github.com/fatih/color"
    "github.com/schollz/progressbar/v3"
)

// Global color variables
var (
    neonPink   = color.New(color.FgHiMagenta, color.Bold)
    neonCyan   = color.New(color.FgHiCyan)
    neonGreen  = color.New(color.FgHiGreen)
    neonYellow = color.New(color.FgHiYellow)
    neonRed    = color.New(color.FgHiRed)
)

// URLProcessor represents the main processor
type URLProcessor struct {
    urlMap       map[string]map[string][]map[string]string
    paramsSeen   map[string]bool
    patternsSeen map[string]bool
    extList      []string
    mu           sync.RWMutex
    stats        ProcessingStats
}

// ProcessingStats tracks processing statistics
type ProcessingStats struct {
    TotalURLs     int
    ValidURLs     int
    InvalidURLs   int
    UniqueHosts   map[string]bool
    UniquePaths   int
    StartTime     time.Time
    ProcessedURLs int64
    mu            sync.Mutex
}

// NewURLProcessor creates a new URL processor instance
func NewURLProcessor() *URLProcessor {
    return &URLProcessor{
        urlMap:       make(map[string]map[string][]map[string]string),
        paramsSeen:   make(map[string]bool),
        patternsSeen: make(map[string]bool),
        extList:      defaultExtensions,
        stats: ProcessingStats{
            UniqueHosts: make(map[string]bool),
            StartTime:   time.Now(),
        },
    }
}

// Default configurations
var defaultExtensions = []string{
    "css", "png", "jpg", "jpeg", "svg", "ico", "webp", "scss",
    "tif", "tiff", "ttf", "otf", "woff", "woff2", "gif",
    "pdf", "bmp", "eot", "mp3", "mp4", "avi",
}

// ProcessURL processes a single URL
func (p *URLProcessor) ProcessURL(rawURL string) {
    p.stats.mu.Lock()
    p.stats.TotalURLs++
    p.stats.mu.Unlock()

    parsedURL, err := url.Parse(rawURL)
    if err != nil {
        p.stats.mu.Lock()
        p.stats.InvalidURLs++
        p.stats.mu.Unlock()
        return
    }

    if parsedURL.Host == "" || parsedURL.Scheme == "" {
        p.stats.mu.Lock()
        p.stats.InvalidURLs++
        p.stats.mu.Unlock()
        return
    }

    host := parsedURL.Scheme + "://" + parsedURL.Host
    path := parsedURL.Path
    query := parsedURL.Query()

    p.mu.Lock()
    defer p.mu.Unlock()

    p.stats.UniqueHosts[host] = true
    p.stats.ValidURLs++

    if _, exists := p.urlMap[host]; !exists {
        p.urlMap[host] = make(map[string][]map[string]string)
    }

    params := make(map[string]string)
    for k, v := range query {
        if len(v) > 0 {
            params[k] = v[0]
            p.paramsSeen[k] = true
        }
    }

    if _, exists := p.urlMap[host][path]; !exists {
        p.urlMap[host][path] = []map[string]string{params}
        p.stats.UniquePaths++
    }
}

// spinnerAnimation displays a loading spinner
func spinnerAnimation(message string, done chan bool) {
    spinners := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
    i := 0
    for {
        select {
        case <-done:
            fmt.Printf("\r%s... Done!%s\n", message, strings.Repeat(" ", 10))
            return
        default:
            fmt.Printf("\r%s %s", spinners[i], message)
            i = (i + 1) % len(spinners)
            time.Sleep(100 * time.Millisecond)
        }
    }
}

func extractDomain(urlStr string) string {
    parsedURL, err := url.Parse(urlStr)
    if err != nil {
        return "output"
    }
    return parsedURL.Hostname()
}

func main() {
    // Command line flags
    inputFile := flag.String("i", "", "Input file containing URLs")
    showHelp := flag.Bool("h", false, "Show help message")
    flag.Parse()

    if *showHelp {
        fmt.Println("Usage: uro [options]")
        fmt.Println("Options:")
        fmt.Println("  -i <file>   Specify input file containing URLs")
        fmt.Println("Example:")
        fmt.Println("  ./uro -i urls.txt")
        fmt.Println("  cat allurls | ./uro")
        return
    }

    processor := NewURLProcessor()

    // Startup banner with developer's name
    neonCyan.Println("╔══════════════════════════════════════════════╗")
    neonCyan.Println("║               URO v1.0.0 Started             ║")
    neonCyan.Println("╚══════════════════════════════════════════════╝")
    neonYellow.Println("             Developed by Karthik-HR0          ")
    
    time.Sleep(500 * time.Millisecond)

    // Input file processing
    var input *os.File
    var err error
    var urls []string
    if *inputFile != "" {
        input, err = os.Open(*inputFile)
        if err != nil {
            neonRed.Printf("[!] Error opening input file: %v\n", err)
            os.Exit(1)
        }
        defer input.Close()
    } else {
        input = os.Stdin
    }

    // Read URLs
    done := make(chan bool)
    go spinnerAnimation("Reading URLs", done)
    
    scanner := bufio.NewScanner(input)
    for scanner.Scan() {
        urls = append(urls, scanner.Text())
    }
    done <- true

    // Process URLs with progress bar
    neonGreen.Printf("\n[+] Found %d URLs\n", len(urls))
    
    bar := progressbar.NewOptions(len(urls),
        progressbar.OptionSetDescription("[+] Processing URLs..."),
        progressbar.OptionSetTheme(progressbar.Theme{
            Saucer:        "=",
            SaucerHead:    ">",
            SaucerPadding: " ",
            BarStart:      "[",
            BarEnd:        "]",
        }))

    for _, url := range urls {
        processor.ProcessURL(url)
        bar.Add(1)
    }

    // Print results
    neonGreen.Printf("\n[+] Processed %d URLs\n", len(urls))
    neonGreen.Printf("[+] Found %d unique hosts\n", len(processor.stats.UniqueHosts))
    neonGreen.Printf("[+] Found %d unique paths\n", processor.stats.UniquePaths)

    // Get the first valid domain from the URLs
    var domain string
    for _, u := range urls {
        if d := extractDomain(u); d != "" {
            domain = d
            break
        }
    }
    if domain == "" {
        domain = "output"
    }

    // Create output file with domain name
    outputFileName := domain + "_uro.txt"
    output, err := os.Create(outputFileName)
    if err != nil {
        neonRed.Printf("\n[!] Error creating output file: %v\n", err)
        os.Exit(1)
    }
    defer output.Close()

    writer := bufio.NewWriter(output)
    for host, paths := range processor.urlMap {
        for path, paramsList := range paths {
            if len(paramsList) == 0 {
                fmt.Fprintln(writer, host+path)
            } else {
                for _, params := range paramsList {
                    queryStr := ""
                    for k, v := range params {
                        if queryStr != "" {
                            queryStr += "&"
                        }
                        queryStr += k + "=" + v
                    }
                    fmt.Fprintln(writer, host+path+"?"+queryStr)
                }
            }
        }
    }
    writer.Flush()

    neonGreen.Printf("[+] Results saved to %s\n", outputFileName)
}
