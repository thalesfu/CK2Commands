package tabuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍尔迈SharmahBarony struct {
	feud.BaseBarony
}

var BSharmah舍尔迈 feud.Barony = &舍尔迈SharmahBarony{}

func init() {
    f := BSharmah舍尔迈.(*舍尔迈SharmahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sharmah",
		TitleName: "舍尔迈",
		TitleCode: "b_sharmah",
	}
}
