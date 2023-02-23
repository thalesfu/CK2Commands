package spoleto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡夏CasciaBarony struct {
	feud.BaseBarony
}

var BCascia卡夏 feud.Barony = &卡夏CasciaBarony{}

func init() {
	f := BCascia卡夏.(*卡夏CasciaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cascia",
		TitleName: "卡夏",
		TitleCode: "b_cascia",
	}
}
