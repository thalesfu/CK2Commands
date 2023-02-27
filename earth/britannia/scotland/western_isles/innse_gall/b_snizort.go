package innse_gall

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯奈泽特SnizortBarony struct {
	feud.BaseBarony
}

var BSnizort斯奈泽特 feud.Barony = &斯奈泽特SnizortBarony{}

func init() {
    f := BSnizort斯奈泽特.(*斯奈泽特SnizortBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "snizort",
		TitleName: "斯奈泽特",
		TitleCode: "b_snizort",
	}
}
