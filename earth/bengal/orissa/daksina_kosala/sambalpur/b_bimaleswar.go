package sambalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗摩丽湿伐罗BimaleswarBarony struct {
	feud.BaseBarony
}

var BBimaleswar毗摩丽湿伐罗 feud.Barony = &毗摩丽湿伐罗BimaleswarBarony{}

func init() {
    f := BBimaleswar毗摩丽湿伐罗.(*毗摩丽湿伐罗BimaleswarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bimaleswar",
		TitleName: "毗摩丽湿伐罗",
		TitleCode: "b_bimaleswar",
	}
}
