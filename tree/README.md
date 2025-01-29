
# Treego CLI App

`treego` is a command-line tool that prints a formatted directory tree with various customizable options. Below are the available functionalities you can use.

## Features
- Print tree structure recursively with proper formatting.
- Print relative path to the directory being searched.
- Print only directories, excluding files.
- Traverse specified nested levels.
- Print file permissions for all files.
- Sort the tree by the last modification time.
- Output the directory tree in XML or JSON format.
- Print indentation lines without excessive whitespace.

## Build and Test Instructions
Use the following `make` commands to build, test, and analyze the project:
- `make build`: Build the `treego` binary.
- `make test`: Run the test suite.
- `make testcover`: Visualize test coverage.

## Usage
### Basic Commands
1. Print Tree Structure Recursively:
   ```
   treego <directory>
   ```
   Example:
   ```
   treego src
   ```

2. Print Relative Path:
   ```
   treego -f <directory>
   ```
   Example:
   ```
   treego -f src
   ```

3. Print Only Directories:
   ```
   treego -d <directory>
   ```
   Example:
   ```
   treego -d src
   ```

4. Traverse Specified Nested Levels:
   ```
   treego -L <level> <directory>
   ```
   Example:
   ```
   treego -L 3 src
   ```

5. Print File Permissions:
   ```
   treego -p <directory>
   ```
   Example:
   ```
   treego -p src
   ```

6. Sort by Last Modification Time:
   ```
   treego -t <directory>
   ```
   Example:
   ```
   treego -t src
   ```

7. Output in XML or JSON Format:
   - XML format:
     ```
     treego -X -L <level> <directory>
     ```
   - JSON format:
     ```
     treego -J -L <level> <directory>
     ```
   Example (XML):
   ```
   treego -X -L 4 src
   ```
   Example (JSON):
   ```
   treego -J -L 4 src
   ```

8. Print Indentation Lines (No Whitespace):
   ```
   treego -if <directory>
   ```
   Example:
   ```
   treego -if src
   ```

### Combining Flags
You can combine multiple flags for more customized searches, just like the actual `tree` command. Here are some examples:

1. **Relative Path with Depth:**
   ```
   treego -f -L 2 src
   ```
   This command will print the relative path of the directories, limiting the output to a depth of 2.

2. **File Permissions with Relative Path:**
   ```
   treego -f -p src
   ```
   This will print the relative path along with file permissions.

3. **Sort by Modification Time with Depth:**
   ```
   treego -t -L 3 src
   ```
   This will sort the tree by the last modification time and limit the output to a depth of 3.

4. **XML Output with Depth and Indentation Lines:**
   ```
   treego -X -L 2 -if src
   ```
   This will output the directory tree in XML format with indentation lines, limiting the depth to 2.

5. **JSON Output with Permissions and Depth:**
   ```
   treego -J -p -L 4 src
   ```
   This will output the directory tree in JSON format with file permissions, limiting the depth to 4.

## Implementation Details
- **Efficient Memory Usage**: Reads files efficiently and processes them without loading entire files into memory.
- **Recursive Search**: Handles nested directories with various options to control depth and details.
- **Flag Management**: Utilizes robust flag handling for customizable output.

---

This project is a simple and efficient implementation of a directory tree listing tool in Go which is built by following One2N Bootcamp - https://one2n.io/go-bootcamp/go-projects/tree-in-go/tree-exercise
