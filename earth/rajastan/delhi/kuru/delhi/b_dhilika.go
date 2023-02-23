package delhi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提梨迦DhilikaBarony struct {
	feud.BaseBarony
}

var BDhilika提梨迦 feud.Barony = &提梨迦DhilikaBarony{}

func init() {
	f := BDhilika提梨迦.(*提梨迦DhilikaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhilika",
		TitleName: "提梨迦",
		TitleCode: "b_dhilika",
	}
}
