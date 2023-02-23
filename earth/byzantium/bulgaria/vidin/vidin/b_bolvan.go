package vidin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔万BolvanBarony struct {
	feud.BaseBarony
}

var BBolvan博尔万 feud.Barony = &博尔万BolvanBarony{}

func init() {
	f := BBolvan博尔万.(*博尔万BolvanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bolvan",
		TitleName: "博尔万",
		TitleCode: "b_bolvan",
	}
}
