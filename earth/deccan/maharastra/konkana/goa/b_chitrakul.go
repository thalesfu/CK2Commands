package goa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 质多罗俱罗ChitrakulBarony struct {
	feud.BaseBarony
}

var BChitrakul质多罗俱罗 feud.Barony = &质多罗俱罗ChitrakulBarony{}

func init() {
	f := BChitrakul质多罗俱罗.(*质多罗俱罗ChitrakulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chitrakul",
		TitleName: "质多罗俱罗",
		TitleCode: "b_chitrakul",
	}
}
