package tyrus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 推罗TyrusBarony struct {
	feud.BaseBarony
}

var BTyrus推罗 feud.Barony = &推罗TyrusBarony{}

func init() {
	f := BTyrus推罗.(*推罗TyrusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tyrus",
		TitleName: "推罗",
		TitleCode: "b_tyrus",
	}
}
