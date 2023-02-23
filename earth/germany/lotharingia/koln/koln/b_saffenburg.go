package koln

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨芬堡SaffenburgBarony struct {
	feud.BaseBarony
}

var BSaffenburg萨芬堡 feud.Barony = &萨芬堡SaffenburgBarony{}

func init() {
	f := BSaffenburg萨芬堡.(*萨芬堡SaffenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saffenburg",
		TitleName: "萨芬堡",
		TitleCode: "b_saffenburg",
	}
}
