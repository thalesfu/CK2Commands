package ulm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茨维法尔滕ZwiefaltenBarony struct {
	feud.BaseBarony
}

var BZwiefalten茨维法尔滕 feud.Barony = &茨维法尔滕ZwiefaltenBarony{}

func init() {
    f := BZwiefalten茨维法尔滕.(*茨维法尔滕ZwiefaltenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zwiefalten",
		TitleName: "茨维法尔滕",
		TitleCode: "b_zwiefalten",
	}
}
