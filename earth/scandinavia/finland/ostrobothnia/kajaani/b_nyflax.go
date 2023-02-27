package kajaani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼夫拉克斯NyflaxBarony struct {
	feud.BaseBarony
}

var BNyflax尼夫拉克斯 feud.Barony = &尼夫拉克斯NyflaxBarony{}

func init() {
    f := BNyflax尼夫拉克斯.(*尼夫拉克斯NyflaxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nyflax",
		TitleName: "尼夫拉克斯",
		TitleCode: "b_nyflax",
	}
}
