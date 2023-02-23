package porto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波尔图PortoBarony struct {
	feud.BaseBarony
}

var BPorto波尔图 feud.Barony = &波尔图PortoBarony{}

func init() {
	f := BPorto波尔图.(*波尔图PortoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "porto",
		TitleName: "波尔图",
		TitleCode: "b_porto",
	}
}
