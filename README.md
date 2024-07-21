# Proxy Checker

A fast, multithreaded CLI tool to verify the status of HTTP and SOCKS proxies.

## Features

- Supports both HTTP and SOCKS proxies
- Multithreaded processing for high-speed checking
- Simple command-line interface
- Outputs results to both console and file

## Installation

1. Ensure you have Go installed on your system (version 1.16 or later recommended).
2. Clone this repository:
   ```
   git clone https://github.com/chetanxpro/proxy-checker.git
   ```
3. Navigate to the project directory:
   ```
   cd proxy-checker
   ```
4. Build the project:
   ```
   go build -o proxy-checker
   ```

## Usage

Run the tool using the following command:

```
./proxy-checker -input <input_file> -output <output_file> -threads <number_of_threads>
```

- `<input_file>`: Path to the file containing the list of proxies (one per line)
- `<output_file>`: Path where the results will be saved
- `<number_of_threads>`: Number of concurrent threads to use for checking

Example:
```
./proxy-checker -input proxies.txt -output results.txt -threads 10
```

## Input File Format

The input file should contain one proxy per line. Supported formats:

- For HTTP proxies: `http://ip:port` or `https://ip:port`
- For SOCKS proxies: `socks5://ip:port` or `socks4://ip:port`

## Output

The tool will display results in real-time on the console and save them to the specified output file. Each line in the output will be marked as either:

- ✅ LIVE: [proxy]
- ❌ DEAD: [proxy]

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
