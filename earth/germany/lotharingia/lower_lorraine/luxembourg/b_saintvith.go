package luxembourg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣维特SaintvithBarony struct {
	feud.BaseBarony
}

var BSaintvith圣维特 feud.Barony = &圣维特SaintvithBarony{}

func init() {
    f := BSaintvith圣维特.(*圣维特SaintvithBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saintvith",
		TitleName: "圣维特",
		TitleCode: "b_saintvith",
	}
}
