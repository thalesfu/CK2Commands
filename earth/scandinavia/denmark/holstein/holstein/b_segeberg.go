package holstein

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞格贝格SegebergBarony struct {
	feud.BaseBarony
}

var BSegeberg塞格贝格 feud.Barony = &塞格贝格SegebergBarony{}

func init() {
	f := BSegeberg塞格贝格.(*塞格贝格SegebergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "segeberg",
		TitleName: "塞格贝格",
		TitleCode: "b_segeberg",
	}
}
