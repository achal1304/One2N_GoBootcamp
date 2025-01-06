# grep_go

`grep_go` is a Go-based implementation of the Linux `grep` command, designed to handle nested files and folders efficiently. It leverages Go's concurrency model using goroutines and channels to maximize performance. The project uses the `cobra` library for managing flags and command-line interface functionality.

## Features
- Search for keywords in files and folders.
- Perform case-insensitive searches.
- Output results to a file (avoids overwriting existing files).
- Search recursively in directories, processing all `.TXT` files.
- Display lines before, after, or both around a matched line.
- Combine multiple flags to customize search behavior.

## Build and Test Instructions
Use the following `make` commands to build, test, and analyze the project:
- `make build`: Build the `grep_go` binary.
- `make test`: Run the test suite.
- `make testcover`: Visualize test coverage.

## Usage
### Basic Commands
1. Search for a keyword from `stdin`:
   ```
   ./grep_go "<search keyword>"
   ```
   Accepts input from `stdin` and processes the search. Output is displayed after receiving an EOF signal (Ctrl+D).

2. Search for a keyword in a specific file:
   ```
   ./grep_go <search keyword> <filename>
   ```

3. Perform a case-insensitive search:
   ```
   ./grep_go -i <search word> <filename>
   ```

4. Output search results to a file:
   ```
   ./grep_go <search word> <filename> -o <outputfilename>
   ```
   **Note**: The output file should not already exist; otherwise, the program will throw an error.

5. Recursively search for a keyword in a folder:
   ```
   ./grep_go -r <searchword> <foldername>
   ```
   Searches recursively in the folder, processing all `.TXT` files.

6. Return `n` lines after, before, or around a search match:
   - Lines after a match:
     ```
     ./grep_go <search word> <filename> -A <count>
     ```
   - Lines before a match:
     ```
     ./grep_go <search word> <filename> -B <count>
     ```
   - Lines before and after a match:
     ```
     ./grep_go <search word> <filename> -C <count>
     ```

### Combining Flags
All flags can be combined for more customized searches. For example:
```
./grep_go -i -r "search word" <foldername> -o <outputfilename> -C 2
```
This command:
- Performs a case-insensitive search.
- Searches recursively in directories.
- Outputs the results to `outputfilename`.
- Includes 2 lines before and after each matched line.

## Implementation Details
- **Concurrency**: Uses goroutines and channels to achieve high performance by parallelizing the search process.
- **Efficient Memory Usage**: Reads files line-by-line instead of loading entire files into memory, avoiding memory exhaustion for large files.
- **Recursive Search**: Handles nested directories and processes all `.TXT` files efficiently.
- **Flag Management**: Utilizes the `cobra` library to handle command-line arguments and flags, making the interface robust and extensible.

## Areas for Improvement
Currently, results are stored in a map until all processing is complete. This keeps results in memory, which can be problematic for very large datasets. A proposed improvement is to utilize channels to print output as soon as a file produces results, reducing memory usage.

---

This project is based on the problem statement from the [One2N Go Bootcamp](https://one2n.io/go-bootcamp/go-projects/grep-in-go/grep-exercise).
