package bearn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔泰兹OrthezBarony struct {
	feud.BaseBarony
}

var BOrthez奥尔泰兹 feud.Barony = &奥尔泰兹OrthezBarony{}

func init() {
    f := BOrthez奥尔泰兹.(*奥尔泰兹OrthezBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orthez",
		TitleName: "奥尔泰兹",
		TitleCode: "b_orthez",
	}
}
