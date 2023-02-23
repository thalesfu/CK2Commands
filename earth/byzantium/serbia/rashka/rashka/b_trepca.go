package rashka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特雷普查TrepcaBarony struct {
	feud.BaseBarony
}

var BTrepca特雷普查 feud.Barony = &特雷普查TrepcaBarony{}

func init() {
	f := BTrepca特雷普查.(*特雷普查TrepcaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trepca",
		TitleName: "特雷普查",
		TitleCode: "b_trepca",
	}
}
