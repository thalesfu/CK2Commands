package bolzano

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔蒂纳CortinaBarony struct {
	feud.BaseBarony
}

var BCortina科尔蒂纳 feud.Barony = &科尔蒂纳CortinaBarony{}

func init() {
	f := BCortina科尔蒂纳.(*科尔蒂纳CortinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cortina",
		TitleName: "科尔蒂纳",
		TitleCode: "b_cortina",
	}
}
