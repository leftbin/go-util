package os

import "github.com/skratchdot/open-golang/open"

func OpenBrowser(url string, dry bool) {
	if url != "" && !dry {
		_ = open.Run(url)
	}
}
