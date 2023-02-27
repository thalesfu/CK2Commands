package semien

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇阿韦塔CiarvetaBarony struct {
	feud.BaseBarony
}

var BCiarveta奇阿韦塔 feud.Barony = &奇阿韦塔CiarvetaBarony{}

func init() {
    f := BCiarveta奇阿韦塔.(*奇阿韦塔CiarvetaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ciarveta",
		TitleName: "奇阿韦塔",
		TitleCode: "b_ciarveta",
	}
}
