package brugge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔登堡AardenburgBarony struct {
	feud.BaseBarony
}

var BAardenburg阿尔登堡 feud.Barony = &阿尔登堡AardenburgBarony{}

func init() {
    f := BAardenburg阿尔登堡.(*阿尔登堡AardenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aardenburg",
		TitleName: "阿尔登堡",
		TitleCode: "b_aardenburg",
	}
}
