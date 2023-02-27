package amorion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多基米翁Docimium_amorionBarony struct {
	feud.BaseBarony
}

var BDocimium_amorion多基米翁 feud.Barony = &多基米翁Docimium_amorionBarony{}

func init() {
    f := BDocimium_amorion多基米翁.(*多基米翁Docimium_amorionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "docimium_amorion",
		TitleName: "多基米翁",
		TitleCode: "b_docimium_amorion",
	}
}
