package varanasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴多希BhadohiBarony struct {
	feud.BaseBarony
}

var BBhadohi巴多希 feud.Barony = &巴多希BhadohiBarony{}

func init() {
	f := BBhadohi巴多希.(*巴多希BhadohiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhadohi",
		TitleName: "巴多希",
		TitleCode: "b_bhadohi",
	}
}
