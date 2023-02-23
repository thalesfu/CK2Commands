package tudgha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾勒尼夫AlnifBarony struct {
	feud.BaseBarony
}

var BAlnif艾勒尼夫 feud.Barony = &艾勒尼夫AlnifBarony{}

func init() {
	f := BAlnif艾勒尼夫.(*艾勒尼夫AlnifBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alnif",
		TitleName: "艾勒尼夫",
		TitleCode: "b_alnif",
	}
}
