# Tidy
Make symlinks from a dotfile dir. Like gnu stow, but use config files instead
```
Tidy Dotfile Manager

Usage:
  tidy [flags]

Examples:
tidy [-s | -u] [directory | *]

Flags:
  -h, --help    help for tidy
  -s, --set     use to create symlinks, mutually exclusive with unset
  -u, --unset   use to remove symlinks, mutually exclusive with set
```
## Config File
```jsonc
{
	"home": [], // ~
	"config": [], // ~/.config
	"bin": [] // ~/.local/bin
}
```
## Example
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
