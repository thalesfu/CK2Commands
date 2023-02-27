package kazakh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔萨克帕伊KarsakpayBarony struct {
	feud.BaseBarony
}

var BKarsakpay卡尔萨克帕伊 feud.Barony = &卡尔萨克帕伊KarsakpayBarony{}

func init() {
    f := BKarsakpay卡尔萨克帕伊.(*卡尔萨克帕伊KarsakpayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karsakpay",
		TitleName: "卡尔萨克帕伊",
		TitleCode: "b_karsakpay",
	}
}
