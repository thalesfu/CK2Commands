package merya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尔日斯克VolzhskBarony struct {
	feud.BaseBarony
}

var BVolzhsk沃尔日斯克 feud.Barony = &沃尔日斯克VolzhskBarony{}

func init() {
    f := BVolzhsk沃尔日斯克.(*沃尔日斯克VolzhskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "volzhsk",
		TitleName: "沃尔日斯克",
		TitleCode: "b_volzhsk",
	}
}
