package tarragona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣库加SancugatBarony struct {
	feud.BaseBarony
}

var BSancugat圣库加 feud.Barony = &圣库加SancugatBarony{}

func init() {
    f := BSancugat圣库加.(*圣库加SancugatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sancugat",
		TitleName: "圣库加",
		TitleCode: "b_sancugat",
	}
}
