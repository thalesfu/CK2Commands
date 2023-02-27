package tinmallal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德尔鲁尔Dar_er_rhoulBarony struct {
	feud.BaseBarony
}

var BDar_er_rhoul德尔鲁尔 feud.Barony = &德尔鲁尔Dar_er_rhoulBarony{}

func init() {
    f := BDar_er_rhoul德尔鲁尔.(*德尔鲁尔Dar_er_rhoulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dar_er_rhoul",
		TitleName: "德尔鲁尔",
		TitleCode: "b_dar_er_rhoul",
	}
}
