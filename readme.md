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
## Install
### Dependancies
- [Go](https://go.dev/)
### Install
```sh
git clone https://github.com/shrimp332/tidy
cd tidy
sudo make install
```
## Config File
```jsonc
{
	"home": [], // ~/
	"config": [], // ~/.config
	"bin": [] // ~/.local/bin
}
```
## Usage
### File Structure
```
dotfiles/
├── scripts/
│  ├── .tidy.json
│  └── increasevol.sh
└── zsh/
   ├── .tidy.json
   ├── .zshrc
   └── zsh/
```
### Config File
```jsonc
// dotfiles/zsh/.tidy.json
{
	"home": [".zshrc"],
	"config": ["zsh"]
}

// dotfiles/scripts/.tidy.json
{
	"bin": ["increasevol.sh"]
}
```
### Commands
```sh
cd dotfiles
tidy -s * # To create both zsh and scripts symlinks
tidy -u scripts # To delete scripts symlink
```
