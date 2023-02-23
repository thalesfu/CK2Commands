package kirghiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帖良古惕TelengitBarony struct {
	feud.BaseBarony
}

var BTelengit帖良古惕 feud.Barony = &帖良古惕TelengitBarony{}

func init() {
	f := BTelengit帖良古惕.(*帖良古惕TelengitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "telengit",
		TitleName: "帖良古惕",
		TitleCode: "b_telengit",
	}
}
