package vitebsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维捷布斯克VitebskBarony struct {
	feud.BaseBarony
}

var BVitebsk维捷布斯克 feud.Barony = &维捷布斯克VitebskBarony{}

func init() {
    f := BVitebsk维捷布斯克.(*维捷布斯克VitebskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vitebsk",
		TitleName: "维捷布斯克",
		TitleCode: "b_vitebsk",
	}
}
