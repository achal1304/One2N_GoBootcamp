# Word Count Utility (`wc`)

This project provides a high-performance word count utility (`wc`) for processing text files. It is designed to handle large files (5GB+) efficiently by leveraging buffered reads and concurrent processing. The utility offers functionality similar to the standard `wc` command on Linux, with added support for multiple files and interactive input mode.

---

## Features

### Count Lines, Words, and Characters:
- Process text files to calculate the number of lines, words, and characters.
- Supports flags:
  - `-l`: Line count
  - `-w`: Word count
  - `-c`: Character count

### Interactive Input Mode:
- Running `./wc` without a file processes user input interactively. Type the text and press `Ctrl+D` to calculate counts.

### Multiple File Support:
- Provide multiple files as input, and the utility processes them concurrently, returning counts for each file.

### Optimized for Large Files:
- Tested with files over 5GB, providing fast responses by:
  - Reading files in 1MB buffers.
  - Splitting data into lines for efficient processing.
  - Launching multiple goroutines to process lines concurrently.
- Avoids loading entire files into memory, preventing memory outages.

---

## Usage
Run the following command from the parent directory:

### Build the Binary
```bash
make buildwc

# Run the Tests
make testwc

# Run tests with coverage
make testwccover
```

Run the binary
```bash
# To get all details like Linecount, wordcount and character counts
./wc <filename>

# Pass flags to get specific details -l for Linecount, -w for Wordcount & -c for CharacterCount
./wc -l -w -c <filename>

# Run without filename to add input in stdin
./wc

# Multiple files
./wc <file1> <file2> <file3>
```
---

## How it works
### Buffered File Reading:
Reads files in 1MB chunks to avoid loading the entire file into memory.
### Line Splitting:
Each buffer is split into lines for processing.
### Concurrent Processing:
Goroutines are launched to process lines concurrently.
Results are aggregated for the total count.
### Efficient Memory Management:
Files are not loaded entirely into memory, ensuring consistent performance and preventing memory exhaustion.
### Multiple File Handling:
Each file is processed in a separate goroutine, allowing fast and parallel execution.

