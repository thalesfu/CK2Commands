package boqtybay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尔德科尔TaldykolBarony struct {
	feud.BaseBarony
}

var BTaldykol塔尔德科尔 feud.Barony = &塔尔德科尔TaldykolBarony{}

func init() {
    f := BTaldykol塔尔德科尔.(*塔尔德科尔TaldykolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taldykol",
		TitleName: "塔尔德科尔",
		TitleCode: "b_taldykol",
	}
}
