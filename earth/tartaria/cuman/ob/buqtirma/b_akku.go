package buqtirma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿库AkkuBarony struct {
	feud.BaseBarony
}

var BAkku阿库 feud.Barony = &阿库AkkuBarony{}

func init() {
    f := BAkku阿库.(*阿库AkkuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akku",
		TitleName: "阿库",
		TitleCode: "b_akku",
	}
}
