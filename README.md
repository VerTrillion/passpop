# Passpop

**Passpop** is a lightweight, secure CLI password manager for macOS and Linux.

## ✨ Features

- 🔐 Store credentials in encrypted YAML files
- 🧪 AES-GCM 256 encryption with secure key
- 🧑‍💻 Easy CLI usage:
  - `passpop init [-s]` — Initialize with or without master password
  - `passpop add -k <key> -p <password>` — Add or update a password
  - `passpop get <key>` — Decrypt and copy password to clipboard
  - `passpop ls` — List all stored keys
  - `passpop rm <key>` — Delete a stored credential
- 🔏 Secure config storage (`~/.passpop/config.yml`)
- 🔄 Auto-export encryption key to `.zshrc` or derive from password
- 💻 Works on macOS and Linux

## 🚀 Installation

Download pre-built binaries from the [GitHub Releases](https://github.com/VerTrillion/passpop/releases) page.

Or build from source:

```bash
git clone https://github.com/VerTrillion/passpop.git
cd passpop
go build -o passpop main.go
```

## 🔧 Usage

```bash
passpop init -s                     # Secure mode with master password
passpop add -k gmail -p secret     # Encrypt and save password
passpop get gmail                  # Decrypt and copy to clipboard
passpop ls                         # List all keys
passpop rm gmail                   # Remove a key
```

## 🔐 Security Considerations

- Credentials are encrypted using AES-GCM 256-bit
- Encryption key is stored in environment variable (`$PASSPOP_KEY`)
- Optionally protect access with a master password
- Files use strict permissions (`chmod 600`)
- Keep your `.zshrc` and machine secure

## 📄 License

MIT © Nuttapong Sudjai
