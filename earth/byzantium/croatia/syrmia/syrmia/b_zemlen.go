package syrmia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泽姆伦ZemlenBarony struct {
	feud.BaseBarony
}

var BZemlen泽姆伦 feud.Barony = &泽姆伦ZemlenBarony{}

func init() {
    f := BZemlen泽姆伦.(*泽姆伦ZemlenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zemlen",
		TitleName: "泽姆伦",
		TitleCode: "b_zemlen",
	}
}
