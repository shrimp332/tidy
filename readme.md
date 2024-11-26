# Tidy
Create symlinks for your dotfiles
```
Tidy Dotfile Linker v1.2.0

Usage:
  tidy [flags]

Examples:
tidy [-s | -u] [directory | *]

Flags:
  -f, --force   overwrite existing files
  -h, --help    help for tidy
  -s, --set     use to create symlinks, mutually exclusive with unset
  -u, --unset   use to remove symlinks, mutually exclusive with set
```
## Installation
### Build from source
#### Dependencies
- [Go](https://go.dev/)
#### Install
```sh
git clone https://github.com/shrimp332/tidy
cd tidy
sudo make install # installs to /usr/local/bin
# or `make install-local` # installs to ~/.local/bin
```
### Using Go
```sh
go install github.com/shrimp332/tidy/cmd/tidy@latest
```
### Archlinux
#### Makepkg
Manual way of making package. Will require manual updates
(it's not in the aur)
```sh
git clone https://github.com/shrimp332/tidy/cmd/tidy
cd tidy
makepkg -si
```
#### [Paru](https://github.com/Morganamilo/paru)
Add this repo as a paru source:  
Add this to `/etc/paru.conf`
```
[TIDY]
Url = https://github.com/shrimp332/tidy
Depth = 1
```
then run
```sh
paru -Sy TIDY/tidy-git
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
