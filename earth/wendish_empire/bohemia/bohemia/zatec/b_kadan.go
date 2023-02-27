package zatec

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡丹KadanBarony struct {
	feud.BaseBarony
}

var BKadan卡丹 feud.Barony = &卡丹KadanBarony{}

func init() {
    f := BKadan卡丹.(*卡丹KadanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kadan",
		TitleName: "卡丹",
		TitleCode: "b_kadan",
	}
}
