package aylik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾里克AylikBarony struct {
	feud.BaseBarony
}

var BAylik艾里克 feud.Barony = &艾里克AylikBarony{}

func init() {
	f := BAylik艾里克.(*艾里克AylikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aylik",
		TitleName: "艾里克",
		TitleCode: "b_aylik",
	}
}
