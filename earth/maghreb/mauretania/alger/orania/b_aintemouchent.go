package orania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾因蒂姆沈特AintemouchentBarony struct {
	feud.BaseBarony
}

var BAintemouchent艾因蒂姆沈特 feud.Barony = &艾因蒂姆沈特AintemouchentBarony{}

func init() {
    f := BAintemouchent艾因蒂姆沈特.(*艾因蒂姆沈特AintemouchentBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aintemouchent",
		TitleName: "艾因蒂姆沈特",
		TitleCode: "b_aintemouchent",
	}
}
