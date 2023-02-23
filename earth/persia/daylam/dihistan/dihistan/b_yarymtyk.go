package dihistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚勒姆特克YarymtykBarony struct {
	feud.BaseBarony
}

var BYarymtyk亚勒姆特克 feud.Barony = &亚勒姆特克YarymtykBarony{}

func init() {
	f := BYarymtyk亚勒姆特克.(*亚勒姆特克YarymtykBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yarymtyk",
		TitleName: "亚勒姆特克",
		TitleCode: "b_yarymtyk",
	}
}
