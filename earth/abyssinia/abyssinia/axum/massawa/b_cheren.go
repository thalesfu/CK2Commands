package massawa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克伦CherenBarony struct {
	feud.BaseBarony
}

var BCheren克伦 feud.Barony = &克伦CherenBarony{}

func init() {
	f := BCheren克伦.(*克伦CherenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cheren",
		TitleName: "克伦",
		TitleCode: "b_cheren",
	}
}
