package kashin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈里齐GoritsyBarony struct {
	feud.BaseBarony
}

var BGoritsy戈里齐 feud.Barony = &戈里齐GoritsyBarony{}

func init() {
	f := BGoritsy戈里齐.(*戈里齐GoritsyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "goritsy",
		TitleName: "戈里齐",
		TitleCode: "b_goritsy",
	}
}
