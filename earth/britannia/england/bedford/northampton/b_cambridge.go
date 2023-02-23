package northampton

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 剑桥CambridgeBarony struct {
	feud.BaseBarony
}

var BCambridge剑桥 feud.Barony = &剑桥CambridgeBarony{}

func init() {
	f := BCambridge剑桥.(*剑桥CambridgeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cambridge",
		TitleName: "剑桥",
		TitleCode: "b_cambridge",
	}
}
