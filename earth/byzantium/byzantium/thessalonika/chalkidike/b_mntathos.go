package chalkidike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿索斯山MntathosBarony struct {
	feud.BaseBarony
}

var BMntathos阿索斯山 feud.Barony = &阿索斯山MntathosBarony{}

func init() {
	f := BMntathos阿索斯山.(*阿索斯山MntathosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mntathos",
		TitleName: "阿索斯山",
		TitleCode: "b_mntathos",
	}
}
