package tunis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈马马特HammametBarony struct {
	feud.BaseBarony
}

var BHammamet哈马马特 feud.Barony = &哈马马特HammametBarony{}

func init() {
	f := BHammamet哈马马特.(*哈马马特HammametBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hammamet",
		TitleName: "哈马马特",
		TitleCode: "b_hammamet",
	}
}
