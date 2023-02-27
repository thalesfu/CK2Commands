package atholl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯特拉塔杜StrathardleBarony struct {
	feud.BaseBarony
}

var BStrathardle斯特拉塔杜 feud.Barony = &斯特拉塔杜StrathardleBarony{}

func init() {
    f := BStrathardle斯特拉塔杜.(*斯特拉塔杜StrathardleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "strathardle",
		TitleName: "斯特拉塔杜",
		TitleCode: "b_strathardle",
	}
}
