# URO - URL Processor

**URO** is a command-line tool for parsing and processing URLs from either a file or standard input. It extracts unique hosts and paths, tracks URL processing stats, and saves the results in an organized output file. URO is designed with simplicity, efficiency, and usability in mind, with real-time progress indicators to improve the user experience.

---

## Features

- **File or Pipe Input**: Accepts URLs from an input file or directly from standard input, supporting flexible usage.
- **Detailed Processing**: Extracts unique hosts and paths, identifies valid/invalid URLs, and processes query parameters.
- **Real-Time Feedback**: Displays a loading spinner while reading and a progress bar during URL processing.
- **Organized Output**: Saves results to a file based on the domain name, providing clear organization and easy retrieval.

---

## Installation

Ensure you have Go installed. Download and install it from [Go's official website](https://golang.org/dl/) if necessary.

Clone this repository and navigate to the project directory:

```bash
git clone <repository-url>
cd uro

Then, build the project:

go build -o uro main.go

This creates an executable named uro.


---

Usage

Basic Commands

1. Process URLs from a file

./uro -i urls.txt


2. Process URLs from standard input (e.g., using cat):

cat allurls | ./uro


3. Display help

./uro -h



Examples

Using an Input File:

./uro -i urls.txt

Using Standard Input with cat:

cat allurls | ./uro

Help and Usage Details:

To display the usage information:

./uro -h

This command shows a help message with usage instructions and examples for quick reference.


---

Output

The processed URLs are saved to a file named <domain>_uro.txt, where <domain> is the first valid domain name found in the input URLs. This file contains each unique host, path, and query parameter combination, organized for easy navigation.


---

Options


---

File Structure

main.go - The primary program file, containing the logic for URL processing, parameter handling, and output generation.

README.md - Documentation covering installation, usage, and feature details.



---

Example Output

If the first domain extracted from the URLs is example.com, the output file will be named example.com_uro.txt and contain unique paths and parameters found during processing, organized as shown below:

https://example.com/path1?param=value
https://example.com/path2
https://example.com/path3?param1=value1&param2=value2


---

License

This project is licensed under the MIT License. See LICENSE for details.


---

Contributions

Contributions are welcome! If you find a bug or have a feature request, please open an issue or submit a pull request. For major changes, please discuss them via an issue first to ensure they align with the project’s goals.


---

Author

Developed by Karthik.


---

This README provides comprehensive setup and usage information, along with clear examples, to help new users get started quickly with uro. Let me know if you need any specific sections or details added!
