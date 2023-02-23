package tribandapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿闵AminBarony struct {
	feud.BaseBarony
}

var BAmin阿闵 feud.Barony = &阿闵AminBarony{}

func init() {
	f := BAmin阿闵.(*阿闵AminBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amin",
		TitleName: "阿闵",
		TitleCode: "b_amin",
	}
}
