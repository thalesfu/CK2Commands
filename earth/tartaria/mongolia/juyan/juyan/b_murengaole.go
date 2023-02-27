package juyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 木仁高勒MurengaoleBarony struct {
	feud.BaseBarony
}

var BMurengaole木仁高勒 feud.Barony = &木仁高勒MurengaoleBarony{}

func init() {
    f := BMurengaole木仁高勒.(*木仁高勒MurengaoleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "murengaole",
		TitleName: "木仁高勒",
		TitleCode: "b_murengaole",
	}
}
