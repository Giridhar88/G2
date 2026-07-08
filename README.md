# G2

An encrypted daily journal for the command line. Entries are stored as individual AES-256-GCM encrypted files, one per day, and edited via Neovim.

## Features

- **AES-256-GCM encryption** with a key derived from your password via scrypt (N=2^15, r=8, p=1)
- **Per-day files** stored at `~/.g2files/DD-MM-YYYY.enc` -- one file per day, no database
- **Authenticated encryption** provides both confidentiality and integrity; tampered files or wrong passwords are detected
- **Neovim integration** with a full decrypt-edit-reencrypt workflow using temporary files

## Usage

```
g2 -new              Create or open today's entry
g2 -open DD-MM-YYYY  Open an entry for a specific date
g2 -open latest      Open today's entry (same as -new)
```

On first run, the `~/.g2files/` directory is created automatically.

## Installation

```bash
git clone <repo-url>
cd g2
go build -o g2 .
```

Requires Go 1.25 or later.

## Dependencies

- [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto) -- scrypt key derivation
- [golang.org/x/term](https://pkg.go.dev/golang.org/x/term) -- secure password input
- Neovim (`nvim`) must be installed and available on `PATH`

## File Format

Each `.enc` file is structured as:

| Offset | Size | Field  |
|--------|------|--------|
| 0      | 16   | Salt   |
| 16     | 12   | Nonce  |
| 28     | var  | Ciphertext (AES-256-GCM) |

## Security Notes

- The decrypted content is written to a temporary file on disk while the editor is open. On multi-user systems, this is a potential exposure window.
- Empty entries (new, unedited files) produce 0-byte plaintext and are handled gracefully.
- No remote sync, backup, or sharing features are included.
