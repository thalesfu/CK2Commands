package oualata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂泽尔特TizertBarony struct {
	feud.BaseBarony
}

var BTizert蒂泽尔特 feud.Barony = &蒂泽尔特TizertBarony{}

func init() {
	f := BTizert蒂泽尔特.(*蒂泽尔特TizertBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tizert",
		TitleName: "蒂泽尔特",
		TitleCode: "b_tizert",
	}
}
