package napoli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波尔蒂奇PorticiBarony struct {
	feud.BaseBarony
}

var BPortici波尔蒂奇 feud.Barony = &波尔蒂奇PorticiBarony{}

func init() {
    f := BPortici波尔蒂奇.(*波尔蒂奇PorticiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "portici",
		TitleName: "波尔蒂奇",
		TitleCode: "b_portici",
	}
}
