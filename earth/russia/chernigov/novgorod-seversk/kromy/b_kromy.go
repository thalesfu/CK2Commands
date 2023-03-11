package kromy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克罗梅KromyBarony struct {
	feud.BaseBarony
}

var BKromy克罗梅 feud.Barony = &克罗梅KromyBarony{}

func init() {
    f := BKromy克罗梅.(*克罗梅KromyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kromy",
		TitleName: "克罗梅",
		TitleCode: "b_kromy",
	}
}
