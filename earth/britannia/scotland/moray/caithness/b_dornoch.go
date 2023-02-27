package caithness

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多诺赫DornochBarony struct {
	feud.BaseBarony
}

var BDornoch多诺赫 feud.Barony = &多诺赫DornochBarony{}

func init() {
    f := BDornoch多诺赫.(*多诺赫DornochBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dornoch",
		TitleName: "多诺赫",
		TitleCode: "b_dornoch",
	}
}
