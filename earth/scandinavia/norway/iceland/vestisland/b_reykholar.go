package vestisland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷克侯拉尔ReykholarBarony struct {
	feud.BaseBarony
}

var BReykholar雷克侯拉尔 feud.Barony = &雷克侯拉尔ReykholarBarony{}

func init() {
	f := BReykholar雷克侯拉尔.(*雷克侯拉尔ReykholarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reykholar",
		TitleName: "雷克侯拉尔",
		TitleCode: "b_reykholar",
	}
}
