package leventina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢加诺LuganoBarony struct {
	feud.BaseBarony
}

var BLugano卢加诺 feud.Barony = &卢加诺LuganoBarony{}

func init() {
	f := BLugano卢加诺.(*卢加诺LuganoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lugano",
		TitleName: "卢加诺",
		TitleCode: "b_lugano",
	}
}
