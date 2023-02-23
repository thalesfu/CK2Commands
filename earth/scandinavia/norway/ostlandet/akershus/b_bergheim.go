package akershus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝格海姆BergheimBarony struct {
	feud.BaseBarony
}

var BBergheim贝格海姆 feud.Barony = &贝格海姆BergheimBarony{}

func init() {
	f := BBergheim贝格海姆.(*贝格海姆BergheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bergheim",
		TitleName: "贝格海姆",
		TitleCode: "b_bergheim",
	}
}
