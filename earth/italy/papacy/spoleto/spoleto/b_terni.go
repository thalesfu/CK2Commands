package spoleto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特尔尼TerniBarony struct {
	feud.BaseBarony
}

var BTerni特尔尼 feud.Barony = &特尔尼TerniBarony{}

func init() {
	f := BTerni特尔尼.(*特尔尼TerniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "terni",
		TitleName: "特尔尼",
		TitleCode: "b_terni",
	}
}
