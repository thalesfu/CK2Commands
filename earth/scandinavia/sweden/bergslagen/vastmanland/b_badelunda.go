package vastmanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴德伦达BadelundaBarony struct {
	feud.BaseBarony
}

var BBadelunda巴德伦达 feud.Barony = &巴德伦达BadelundaBarony{}

func init() {
    f := BBadelunda巴德伦达.(*巴德伦达BadelundaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "badelunda",
		TitleName: "巴德伦达",
		TitleCode: "b_badelunda",
	}
}
