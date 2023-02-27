package yeruslan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶鲁斯兰YeruslanBarony struct {
	feud.BaseBarony
}

var BYeruslan叶鲁斯兰 feud.Barony = &叶鲁斯兰YeruslanBarony{}

func init() {
    f := BYeruslan叶鲁斯兰.(*叶鲁斯兰YeruslanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yeruslan",
		TitleName: "叶鲁斯兰",
		TitleCode: "b_yeruslan",
	}
}
