package sarangpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阇罗婆罗JhalawarBarony struct {
	feud.BaseBarony
}

var BJhalawar阇罗婆罗 feud.Barony = &阇罗婆罗JhalawarBarony{}

func init() {
    f := BJhalawar阇罗婆罗.(*阇罗婆罗JhalawarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jhalawar",
		TitleName: "阇罗婆罗",
		TitleCode: "b_jhalawar",
	}
}
