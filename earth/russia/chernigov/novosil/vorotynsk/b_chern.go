package vorotynsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔尼ChernBarony struct {
	feud.BaseBarony
}

var BChern切尔尼 feud.Barony = &切尔尼ChernBarony{}

func init() {
	f := BChern切尔尼.(*切尔尼ChernBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chern",
		TitleName: "切尔尼",
		TitleCode: "b_chern",
	}
}
