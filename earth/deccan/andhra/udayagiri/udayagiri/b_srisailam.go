package udayagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 室利势罗SrisailamBarony struct {
	feud.BaseBarony
}

var BSrisailam室利势罗 feud.Barony = &室利势罗SrisailamBarony{}

func init() {
    f := BSrisailam室利势罗.(*室利势罗SrisailamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "srisailam",
		TitleName: "室利势罗",
		TitleCode: "b_srisailam",
	}
}
