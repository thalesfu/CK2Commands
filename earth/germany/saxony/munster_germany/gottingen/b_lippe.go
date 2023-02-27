package gottingen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利珀LippeBarony struct {
	feud.BaseBarony
}

var BLippe利珀 feud.Barony = &利珀LippeBarony{}

func init() {
    f := BLippe利珀.(*利珀LippeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lippe",
		TitleName: "利珀",
		TitleCode: "b_lippe",
	}
}
