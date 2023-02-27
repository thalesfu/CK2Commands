package aksu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 温宿WensuBarony struct {
	feud.BaseBarony
}

var BWensu温宿 feud.Barony = &温宿WensuBarony{}

func init() {
    f := BWensu温宿.(*温宿WensuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wensu",
		TitleName: "温宿",
		TitleCode: "b_wensu",
	}
}
