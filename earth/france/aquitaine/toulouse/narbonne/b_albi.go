package narbonne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔比AlbiBarony struct {
	feud.BaseBarony
}

var BAlbi阿尔比 feud.Barony = &阿尔比AlbiBarony{}

func init() {
    f := BAlbi阿尔比.(*阿尔比AlbiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "albi",
		TitleName: "阿尔比",
		TitleCode: "b_albi",
	}
}
