package alto_aragon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈卡JacaBarony struct {
	feud.BaseBarony
}

var BJaca哈卡 feud.Barony = &哈卡JacaBarony{}

func init() {
    f := BJaca哈卡.(*哈卡JacaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jaca",
		TitleName: "哈卡",
		TitleCode: "b_jaca",
	}
}
