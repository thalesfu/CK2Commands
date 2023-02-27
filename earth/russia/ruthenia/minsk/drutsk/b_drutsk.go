package drutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德鲁茨克DrutskBarony struct {
	feud.BaseBarony
}

var BDrutsk德鲁茨克 feud.Barony = &德鲁茨克DrutskBarony{}

func init() {
    f := BDrutsk德鲁茨克.(*德鲁茨克DrutskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drutsk",
		TitleName: "德鲁茨克",
		TitleCode: "b_drutsk",
	}
}
