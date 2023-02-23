package volkovysk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尔科维斯克VolkovyskBarony struct {
	feud.BaseBarony
}

var BVolkovysk沃尔科维斯克 feud.Barony = &沃尔科维斯克VolkovyskBarony{}

func init() {
	f := BVolkovysk沃尔科维斯克.(*沃尔科维斯克VolkovyskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "volkovysk",
		TitleName: "沃尔科维斯克",
		TitleCode: "b_volkovysk",
	}
}
