package rostock

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗斯托克RostockBarony struct {
	feud.BaseBarony
}

var BRostock罗斯托克 feud.Barony = &罗斯托克RostockBarony{}

func init() {
    f := BRostock罗斯托克.(*罗斯托克RostockBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rostock",
		TitleName: "罗斯托克",
		TitleCode: "b_rostock",
	}
}
