# URO (URL Resource Organizer)

URO is a command-line tool that processes a list of URLs, extracts parameters, and organizes them into a structured output format. It is designed for developers, security researchers, and anyone who needs to manage and analyze URLs efficiently.

![uro png](https://github.com/user-attachments/assets/98f34543-abd2-4e84-9b20-4c40bfd1f8f0)


## Features

- Processes URLs from a file or standard input.
- Extracts query parameters and organizes them by unique hosts and paths.
- Outputs results to a structured text file.
- Displays processing statistics and a visual progress bar.

## Installation

1. **Clone the repository**:
   ```bash
   go intall https://github.com/Karthik-HR0/uro@latest
   ```
   ### Basic Usage
The quickest way to include uro in your workflow is to feed it data through stdin and print it to your terminal.
```
cat urls.txt | uro
```
### Example output 
```
$ ./uro -i urls.txt
╔══════════════════════════════════════════════╗
║               URO v1.0.0 Started                     ║
╚══════════════════════════════════════════════╝
             Developed by Karthik-HR0          

[+] Found 1000 URLs
[+] Processing URLs... [================>] 100%
[+] Processed 1000 URLs
[+] Found 5 unique hosts
[+] Found 150 unique paths
[+] Results saved to example.com_uro.txt
```
