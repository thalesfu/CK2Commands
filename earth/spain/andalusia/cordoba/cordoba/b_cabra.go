package cordoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡夫拉CabraBarony struct {
	feud.BaseBarony
}

var BCabra卡夫拉 feud.Barony = &卡夫拉CabraBarony{}

func init() {
	f := BCabra卡夫拉.(*卡夫拉CabraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cabra",
		TitleName: "卡夫拉",
		TitleCode: "b_cabra",
	}
}
