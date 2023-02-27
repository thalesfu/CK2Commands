package krems

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍恩Horn_kremsBarony struct {
	feud.BaseBarony
}

var BHorn_krems霍恩 feud.Barony = &霍恩Horn_kremsBarony{}

func init() {
    f := BHorn_krems霍恩.(*霍恩Horn_kremsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "horn_krems",
		TitleName: "霍恩",
		TitleCode: "b_horn_krems",
	}
}
