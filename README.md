# dupl — Fast Duplicate File Finder

`dupl` is a simple command-line tool for scanning a directory and identifying duplicate files based on filename and size.

---

## 🧩 Features

- Recursively scans all files under a given directory
- Detects potential duplicates using filename + size as signature
- Outputs groups of files that appear to be duplicates

---

## 🚀 Usage

```sh
dupl -path /your/target/directory
```

### Short Flags

```sh
dupl -p /your/target/directory
```

If no path is provided, it defaults to the current directory (`.`).

---

## ⚙️ Command-Line Flags

| Flag        | Description                                          |
|-------------|------------------------------------------------------|
| `-path`     | Root directory to scan (default: current directory)  |
| `-p`        | Alias for `-path`                                    |
| `-h`        | Show help information                                |
| `-help`     | Show help information (long form)                    |

---

## 📦 Installation

### Build from Source

1. Clone the repository:
   ```sh
   git clone https://github.com/yourname/dupl.git
   cd dupl
   ```

2. Build the binary:
   ```sh
    go build -o dupl ./cmd/dupl
   ```

3. Run:
   ```sh
   ./dupl -path /your/target

   ./dupl -p /your/target

   ./dupl /your/target
   ```

---

## 🔍 How It Works

Files are grouped by a lightweight signature composed of:

- File **name**
- File **size**

If two or more files share the same signature, they are reported as duplicates.

> Note: This method is fast but may report false positives. Hash-based matching may be added in the future for content verification.

---

## 🛠 Logging

This utility uses a custom `logx` package to handle:

- `Info()` — log general actions
- `Error()` — log recoverable issues
- `FatalErr(err)` — log fatal errors and exit

---

## 🧪 Example Output

```text
Searching for duplicate files in: /photos/archive

INFO  2025/06/12 21:43:22 Duplicates found:
photos/archive/cat.jpg
photos/backup/cat.jpg

INFO  2025/06/12 21:43:22 Duplicates found:
docs/report.pdf
old_docs/copy_of_report.pdf
```

---

## 🧱 Project Structure

```text
dupl/
├── cmd/               # Main CLI entrypoint
│   └── main.go
├── internal/
│   ├── finder/        # Duplicate finding logic
│   │   └── finder.go
│   └── logx/          # Lightweight logging abstraction
│       └── logx.go
├── go.mod
├── LICENSE
└── README.md
```

---

## 📄 License

This project is licensed under the MIT License. See [`LICENSE`](LICENSE) for details.

---


