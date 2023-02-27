package gobir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈比尔GobirBarony struct {
	feud.BaseBarony
}

var BGobir戈比尔 feud.Barony = &戈比尔GobirBarony{}

func init() {
    f := BGobir戈比尔.(*戈比尔GobirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gobir",
		TitleName: "戈比尔",
		TitleCode: "b_gobir",
	}
}
