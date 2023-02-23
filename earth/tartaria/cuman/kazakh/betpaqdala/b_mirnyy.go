package betpaqdala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米尔内MirnyyBarony struct {
	feud.BaseBarony
}

var BMirnyy米尔内 feud.Barony = &米尔内MirnyyBarony{}

func init() {
	f := BMirnyy米尔内.(*米尔内MirnyyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mirnyy",
		TitleName: "米尔内",
		TitleCode: "b_mirnyy",
	}
}
