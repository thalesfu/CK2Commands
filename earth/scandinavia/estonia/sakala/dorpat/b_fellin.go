package dorpat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费林FellinBarony struct {
	feud.BaseBarony
}

var BFellin费林 feud.Barony = &费林FellinBarony{}

func init() {
    f := BFellin费林.(*费林FellinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fellin",
		TitleName: "费林",
		TitleCode: "b_fellin",
	}
}
