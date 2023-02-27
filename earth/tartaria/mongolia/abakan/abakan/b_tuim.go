package abakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图伊姆TuimBarony struct {
	feud.BaseBarony
}

var BTuim图伊姆 feud.Barony = &图伊姆TuimBarony{}

func init() {
    f := BTuim图伊姆.(*图伊姆TuimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tuim",
		TitleName: "图伊姆",
		TitleCode: "b_tuim",
	}
}
