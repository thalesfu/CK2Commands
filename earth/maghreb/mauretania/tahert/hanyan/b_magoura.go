package hanyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马古拉MagouraBarony struct {
	feud.BaseBarony
}

var BMagoura马古拉 feud.Barony = &马古拉MagouraBarony{}

func init() {
	f := BMagoura马古拉.(*马古拉MagouraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "magoura",
		TitleName: "马古拉",
		TitleCode: "b_magoura",
	}
}
