package mangyul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加加GyagyaBarony struct {
	feud.BaseBarony
}

var BGyagya加加 feud.Barony = &加加GyagyaBarony{}

func init() {
    f := BGyagya加加.(*加加GyagyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gyagya",
		TitleName: "加加",
		TitleCode: "b_gyagya",
	}
}
