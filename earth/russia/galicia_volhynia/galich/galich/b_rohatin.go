package galich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗加京RohatinBarony struct {
	feud.BaseBarony
}

var BRohatin罗加京 feud.Barony = &罗加京RohatinBarony{}

func init() {
    f := BRohatin罗加京.(*罗加京RohatinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rohatin",
		TitleName: "罗加京",
		TitleCode: "b_rohatin",
	}
}
