# Tidy
Create symlinks for your dotfiles
```
Tidy Dotfile Linker v2.0.0

Usage:
  tidy [flags]
  tidy [command]

Examples:
tidy [set | unset] [directory | *]

Available Commands:
  help        Help about any command
  set         Create symlinks
  unset       Delete symlinks

Flags:
  -h, --help   help for tidy
```
## Installation
### Build from source
#### Dependencies
- [Go](https://go.dev/)
#### Install
```sh
git clone https://github.com/shrimp332/tidy
cd tidy
sudo make install
```
`make install PREFIX=~/.local` for a user install (assumes ~/.local/bin is in your PATH)
### Using Go
```sh
go install github.com/shrimp332/tidy@latest
```
### Archlinux
#### Makepkg
Manual way of making package. Will require manual updates
(it's not in the aur)
```sh
git clone https://github.com/shrimp332/shrimp-ur
cd shrimp-ur/tidy-git
makepkg -si
```
#### [Paru](https://github.com/Morganamilo/paru)
Add [this](https://github.com/shrimp332/shrimp-ur) as a paru repo. ([instructions](https://github.com/shrimp332/shrimp-ur/blob/main/readme.md))
```sh
paru -S tidy-git
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
tidy set * # To create zsh, scripts, and other symlinks
tidy set -f * #  To force create symlinks (will overwrite existing files)
tidy unset scripts # To delete scripts symlink
```
