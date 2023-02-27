package nilagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩醯湿伐蓝MaheshwaramBarony struct {
	feud.BaseBarony
}

var BMaheshwaram摩醯湿伐蓝 feud.Barony = &摩醯湿伐蓝MaheshwaramBarony{}

func init() {
    f := BMaheshwaram摩醯湿伐蓝.(*摩醯湿伐蓝MaheshwaramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maheshwaram",
		TitleName: "摩醯湿伐蓝",
		TitleCode: "b_maheshwaram",
	}
}
