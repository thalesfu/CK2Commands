package nubia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿纳格AnagBarony struct {
	feud.BaseBarony
}

var BAnag阿纳格 feud.Barony = &阿纳格AnagBarony{}

func init() {
    f := BAnag阿纳格.(*阿纳格AnagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anag",
		TitleName: "阿纳格",
		TitleCode: "b_anag",
	}
}
