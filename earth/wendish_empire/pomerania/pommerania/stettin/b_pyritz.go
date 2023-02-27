package stettin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮里茨PyritzBarony struct {
	feud.BaseBarony
}

var BPyritz皮里茨 feud.Barony = &皮里茨PyritzBarony{}

func init() {
    f := BPyritz皮里茨.(*皮里茨PyritzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pyritz",
		TitleName: "皮里茨",
		TitleCode: "b_pyritz",
	}
}
