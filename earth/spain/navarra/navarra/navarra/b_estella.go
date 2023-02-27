package navarra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯特利亚EstellaBarony struct {
	feud.BaseBarony
}

var BEstella埃斯特利亚 feud.Barony = &埃斯特利亚EstellaBarony{}

func init() {
    f := BEstella埃斯特利亚.(*埃斯特利亚EstellaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "estella",
		TitleName: "埃斯特利亚",
		TitleCode: "b_estella",
	}
}
