package marmaros

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 劳霍RahoBarony struct {
	feud.BaseBarony
}

var BRaho劳霍 feud.Barony = &劳霍RahoBarony{}

func init() {
	f := BRaho劳霍.(*劳霍RahoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raho",
		TitleName: "劳霍",
		TitleCode: "b_raho",
	}
}
