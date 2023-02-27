package corfu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯刻里亚ScheriaBarony struct {
	feud.BaseBarony
}

var BScheria斯刻里亚 feud.Barony = &斯刻里亚ScheriaBarony{}

func init() {
    f := BScheria斯刻里亚.(*斯刻里亚ScheriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "scheria",
		TitleName: "斯刻里亚",
		TitleCode: "b_scheria",
	}
}
