package limisol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿格里迪AgridiBarony struct {
	feud.BaseBarony
}

var BAgridi阿格里迪 feud.Barony = &阿格里迪AgridiBarony{}

func init() {
	f := BAgridi阿格里迪.(*阿格里迪AgridiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agridi",
		TitleName: "阿格里迪",
		TitleCode: "b_agridi",
	}
}
