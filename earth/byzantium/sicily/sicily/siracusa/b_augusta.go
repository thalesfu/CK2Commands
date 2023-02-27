package siracusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥古斯塔AugustaBarony struct {
	feud.BaseBarony
}

var BAugusta奥古斯塔 feud.Barony = &奥古斯塔AugustaBarony{}

func init() {
    f := BAugusta奥古斯塔.(*奥古斯塔AugustaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "augusta",
		TitleName: "奥古斯塔",
		TitleCode: "b_augusta",
	}
}
