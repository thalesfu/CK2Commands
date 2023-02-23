package tirol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣安东StantonBarony struct {
	feud.BaseBarony
}

var BStanton圣安东 feud.Barony = &圣安东StantonBarony{}

func init() {
	f := BStanton圣安东.(*圣安东StantonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stanton",
		TitleName: "圣安东",
		TitleCode: "b_stanton",
	}
}
