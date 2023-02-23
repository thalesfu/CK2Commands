package pavia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮亚琴察PiacenzaBarony struct {
	feud.BaseBarony
}

var BPiacenza皮亚琴察 feud.Barony = &皮亚琴察PiacenzaBarony{}

func init() {
	f := BPiacenza皮亚琴察.(*皮亚琴察PiacenzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "piacenza",
		TitleName: "皮亚琴察",
		TitleCode: "b_piacenza",
	}
}
