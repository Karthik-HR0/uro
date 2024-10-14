# uro (URL Processor)

uro is a lightweight command-line tool written in Go for normalizing and processing URLs. It removes duplicates, sorts query parameters, and helps clean up URL data for tasks such as web scraping, API analysis, and debugging.

# Features
![68747470733a2f2f692e6962622e636f2f783274574343352f75726f2d64656d6f2e706e67](https://github.com/user-attachments/assets/dcecd8e8-5728-4671-ba4c-9f195856c9dd)


Normalize and clean URLs.

Sort query parameters alphabetically.

Remove duplicate URLs.

Handle large lists of URLs efficiently from input files.


Installation

To install goro from the GitHub repository, make sure you have Go installed, and then run the following command:
```
go install github.com/Karthik-HR0/uro/cmd/uro@latest
```

This command will install goro globally on your system as a command-line tool.

# Usage

Once installed, you can use the goro command to process URLs from an input file and write the cleaned output to another file:
```
uro -i input.txt -o output.txt

-i input.txt: Specifies the input file containing a list of URLs (one URL per line).

-o output.txt: Specifies the output file where the normalized URLs will be saved.
```

# Example
```
uro -i urls.txt -o cleaned_urls.txt

If the file urls.txt contains:

https://example.com/path?query=1&action=edit
https://example.com/path?query=2&action=edit
https://example.com/another-path?query=1
https://example.com/path?query=1&action=edit
https://anotherexample.com/path?query=3
```
```
The uro tool will output to cleaned_urls.txt:

https://example.com/path?action=edit&query=1
https://example.com/path?action=edit&query=2
https://example.com/another-path?query=1
https://anotherexample.com/path?query=3
```
# Prerequisites

Go 1.20 or later must be installed. You can download Go from the official Go website.
https://golang.org/dl/


You can verify your Go installation by running:

go version 

# Project Structure

The project follows a clean, organized structure to separate the application logic from utility and test code:
```
uro/
├── cmd/
│   └── uro/
│       └── main.go        # Entry point of the application
├── internal/
│   └── processor/
│       └── url_processor.go # URL processing logic
├── pkg/
│   └── utils/
│       └── helpers.go     # Utility functions
├── test/                  # Test files (optional)
└── go.mod                 # Go module file
```
