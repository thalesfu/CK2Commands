package aydhab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙腊腾ShalateenBarony struct {
	feud.BaseBarony
}

var BShalateen沙腊腾 feud.Barony = &沙腊腾ShalateenBarony{}

func init() {
    f := BShalateen沙腊腾.(*沙腊腾ShalateenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shalateen",
		TitleName: "沙腊腾",
		TitleCode: "b_shalateen",
	}
}
