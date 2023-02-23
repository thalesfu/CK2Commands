package cornouaille

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜瓦讷内DouarnenezBarony struct {
	feud.BaseBarony
}

var BDouarnenez杜瓦讷内 feud.Barony = &杜瓦讷内DouarnenezBarony{}

func init() {
	f := BDouarnenez杜瓦讷内.(*杜瓦讷内DouarnenezBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "douarnenez",
		TitleName: "杜瓦讷内",
		TitleCode: "b_douarnenez",
	}
}
