package coqen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 门东MaindongBarony struct {
	feud.BaseBarony
}

var BMaindong门东 feud.Barony = &门东MaindongBarony{}

func init() {
	f := BMaindong门东.(*门东MaindongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maindong",
		TitleName: "门东",
		TitleCode: "b_maindong",
	}
}
