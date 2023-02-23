package ephesos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷贝多斯LebedosBarony struct {
	feud.BaseBarony
}

var BLebedos雷贝多斯 feud.Barony = &雷贝多斯LebedosBarony{}

func init() {
	f := BLebedos雷贝多斯.(*雷贝多斯LebedosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lebedos",
		TitleName: "雷贝多斯",
		TitleCode: "b_lebedos",
	}
}
