package cremona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索斯皮罗SospiroBarony struct {
	feud.BaseBarony
}

var BSospiro索斯皮罗 feud.Barony = &索斯皮罗SospiroBarony{}

func init() {
    f := BSospiro索斯皮罗.(*索斯皮罗SospiroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sospiro",
		TitleName: "索斯皮罗",
		TitleCode: "b_sospiro",
	}
}
