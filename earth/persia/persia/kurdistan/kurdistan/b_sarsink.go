package kurdistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉斯尼克SarsinkBarony struct {
	feud.BaseBarony
}

var BSarsink萨拉斯尼克 feud.Barony = &萨拉斯尼克SarsinkBarony{}

func init() {
    f := BSarsink萨拉斯尼克.(*萨拉斯尼克SarsinkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarsink",
		TitleName: "萨拉斯尼克",
		TitleCode: "b_sarsink",
	}
}
