# Tidy
Create symlinks for your dotfiles
```
Tidy Dotfile Linker

Usage:
  tidy [flags]

Examples:
tidy [-s | -u] [directory | *]

Flags:
  -h, --help    help for tidy
  -s, --set     use to create symlinks, mutually exclusive with unset
  -u, --unset   use to remove symlinks, mutually exclusive with set
```
## Installation
### Dependencies
- [Go](https://go.dev/)
### Build from source
```sh
git clone https://github.com/shrimp332/tidy
cd tidy
sudo make install # installs to /usr/local/bin
# or `make install-local` # installs to ~/.local/bin
```
## Config File
```jsonc
{
  "home": [], // ~/
  "config": [], // ~/.config
  "bin": [], // ~/.local/bin
  "custom": {
    "directory": [] // custom location
  }
}
```
## Usage
### File Structure
```
dotfiles/
├── scripts/
│  ├── .tidy.json
│  └── increasevol.sh
├── zsh/
│  ├── .tidy.json
│  ├── .zshrc
│  ├── .zshenv
│  └── zsh/
└── other/
   ├── .tidy.json
   ├── passwords.kdbx
   └── obsidian/
```
### Config File
```jsonc
// dotfiles/zsh/.tidy.json
{
  "home": [".zshrc", ".zshenv"],
  "config": ["zsh"]
}

// dotfiles/scripts/.tidy.json
{
  "bin": ["increasevol.sh"]
}

// dotfiles/other/.tidy.json
{
  "custom": {
    "~/Notes": ["obsidian"],
    "~/Documents": ["passwords.kbdx"]
  }
}

```
### Commands
```sh
cd dotfiles
tidy -s * # To create zsh, scripts, and other symlinks
tidy -u scripts # To delete scripts symlink
```
