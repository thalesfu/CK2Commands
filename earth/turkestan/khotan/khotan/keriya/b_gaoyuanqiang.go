package keriya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 高院墙GaoyuanqiangBarony struct {
	feud.BaseBarony
}

var BGaoyuanqiang高院墙 feud.Barony = &高院墙GaoyuanqiangBarony{}

func init() {
    f := BGaoyuanqiang高院墙.(*高院墙GaoyuanqiangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gaoyuanqiang",
		TitleName: "高院墙",
		TitleCode: "b_gaoyuanqiang",
	}
}
