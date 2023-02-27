package coimbra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿威罗AveiroBarony struct {
	feud.BaseBarony
}

var BAveiro阿威罗 feud.Barony = &阿威罗AveiroBarony{}

func init() {
    f := BAveiro阿威罗.(*阿威罗AveiroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aveiro",
		TitleName: "阿威罗",
		TitleCode: "b_aveiro",
	}
}
