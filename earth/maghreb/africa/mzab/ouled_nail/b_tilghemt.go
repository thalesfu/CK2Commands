package ouled_nail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提勒盖姆特TilghemtBarony struct {
	feud.BaseBarony
}

var BTilghemt提勒盖姆特 feud.Barony = &提勒盖姆特TilghemtBarony{}

func init() {
    f := BTilghemt提勒盖姆特.(*提勒盖姆特TilghemtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tilghemt",
		TitleName: "提勒盖姆特",
		TitleCode: "b_tilghemt",
	}
}
