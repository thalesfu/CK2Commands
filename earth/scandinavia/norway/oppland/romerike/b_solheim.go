package romerike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索尔海姆SolheimBarony struct {
	feud.BaseBarony
}

var BSolheim索尔海姆 feud.Barony = &索尔海姆SolheimBarony{}

func init() {
	f := BSolheim索尔海姆.(*索尔海姆SolheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "solheim",
		TitleName: "索尔海姆",
		TitleCode: "b_solheim",
	}
}
