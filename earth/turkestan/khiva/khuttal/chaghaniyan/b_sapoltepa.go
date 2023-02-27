package chaghaniyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨波特帕SapoltepaBarony struct {
	feud.BaseBarony
}

var BSapoltepa萨波特帕 feud.Barony = &萨波特帕SapoltepaBarony{}

func init() {
    f := BSapoltepa萨波特帕.(*萨波特帕SapoltepaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sapoltepa",
		TitleName: "萨波特帕",
		TitleCode: "b_sapoltepa",
	}
}
