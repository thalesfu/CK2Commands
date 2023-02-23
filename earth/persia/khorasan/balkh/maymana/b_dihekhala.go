package maymana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪厄哈拉DihekhalaBarony struct {
	feud.BaseBarony
}

var BDihekhala迪厄哈拉 feud.Barony = &迪厄哈拉DihekhalaBarony{}

func init() {
	f := BDihekhala迪厄哈拉.(*迪厄哈拉DihekhalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dihekhala",
		TitleName: "迪厄哈拉",
		TitleCode: "b_dihekhala",
	}
}
