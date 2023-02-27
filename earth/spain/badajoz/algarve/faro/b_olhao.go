package faro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥良OlhaoBarony struct {
	feud.BaseBarony
}

var BOlhao奥良 feud.Barony = &奥良OlhaoBarony{}

func init() {
    f := BOlhao奥良.(*奥良OlhaoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olhao",
		TitleName: "奥良",
		TitleCode: "b_olhao",
	}
}
