package guines

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣奥梅尔SaintomerBarony struct {
	feud.BaseBarony
}

var BSaintomer圣奥梅尔 feud.Barony = &圣奥梅尔SaintomerBarony{}

func init() {
	f := BSaintomer圣奥梅尔.(*圣奥梅尔SaintomerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saintomer",
		TitleName: "圣奥梅尔",
		TitleCode: "b_saintomer",
	}
}
