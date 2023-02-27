package prusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多西米翁DocimiumBarony struct {
	feud.BaseBarony
}

var BDocimium多西米翁 feud.Barony = &多西米翁DocimiumBarony{}

func init() {
    f := BDocimium多西米翁.(*多西米翁DocimiumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "docimium",
		TitleName: "多西米翁",
		TitleCode: "b_docimium",
	}
}
