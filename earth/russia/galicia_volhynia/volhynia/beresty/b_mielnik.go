package beresty

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅尔尼克MielnikBarony struct {
	feud.BaseBarony
}

var BMielnik梅尔尼克 feud.Barony = &梅尔尼克MielnikBarony{}

func init() {
    f := BMielnik梅尔尼克.(*梅尔尼克MielnikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mielnik",
		TitleName: "梅尔尼克",
		TitleCode: "b_mielnik",
	}
}
