package djenne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德拉里DeraryBarony struct {
	feud.BaseBarony
}

var BDerary德拉里 feud.Barony = &德拉里DeraryBarony{}

func init() {
    f := BDerary德拉里.(*德拉里DeraryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "derary",
		TitleName: "德拉里",
		TitleCode: "b_derary",
	}
}
