package tummana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 道先TorxemBarony struct {
	feud.BaseBarony
}

var BTorxem道先 feud.Barony = &道先TorxemBarony{}

func init() {
	f := BTorxem道先.(*道先TorxemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "torxem",
		TitleName: "道先",
		TitleCode: "b_torxem",
	}
}
