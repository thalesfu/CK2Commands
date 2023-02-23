package asni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内村BhitargaonBarony struct {
	feud.BaseBarony
}

var BBhitargaon内村 feud.Barony = &内村BhitargaonBarony{}

func init() {
	f := BBhitargaon内村.(*内村BhitargaonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhitargaon",
		TitleName: "内村",
		TitleCode: "b_bhitargaon",
	}
}
