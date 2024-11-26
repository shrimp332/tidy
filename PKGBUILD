pkgname=tidy-git
_pkgname=tidy
pkgver=1.2.0
pkgrel=1
pkgdesc='dotfile symlinker'
arch=('x86_64')
url="https://github.com/shrimp332/tidy"
license=('BSD-3-Clause')
makedepends=('git' 'go>=1.23' 'make')
options=(!debug)
source=("git+https://github.com/shrimp332/tidy.git")
md5sums=('SKIP')

pkgver() {
	cd "$_pkgname"
	git describe --tags | sed 's/-\([0-9]\)-/.r\1./;s/v//'
}

build() {
	cd "$_pkgname"
	make DESTDIR="$pkgdir/" build
}

package() {
	cd "$_pkgname"
	install -Dm755 bin/"$_pkgname" "$pkgdir"/usr/bin/tidy
}
