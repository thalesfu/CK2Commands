package tyrus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞达洛斯ScandalonBarony struct {
	feud.BaseBarony
}

var BScandalon塞达洛斯 feud.Barony = &塞达洛斯ScandalonBarony{}

func init() {
    f := BScandalon塞达洛斯.(*塞达洛斯ScandalonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "scandalon",
		TitleName: "塞达洛斯",
		TitleCode: "b_scandalon",
	}
}
