package sijilmasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅尔祖卡MerzougaBarony struct {
	feud.BaseBarony
}

var BMerzouga梅尔祖卡 feud.Barony = &梅尔祖卡MerzougaBarony{}

func init() {
    f := BMerzouga梅尔祖卡.(*梅尔祖卡MerzougaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "merzouga",
		TitleName: "梅尔祖卡",
		TitleCode: "b_merzouga",
	}
}
