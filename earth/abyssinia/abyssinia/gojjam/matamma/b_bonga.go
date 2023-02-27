package matamma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邦加BongaBarony struct {
	feud.BaseBarony
}

var BBonga邦加 feud.Barony = &邦加BongaBarony{}

func init() {
    f := BBonga邦加.(*邦加BongaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bonga",
		TitleName: "邦加",
		TitleCode: "b_bonga",
	}
}
