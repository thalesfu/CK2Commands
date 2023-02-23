package romsdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 许斯塔HustadBarony struct {
	feud.BaseBarony
}

var BHustad许斯塔 feud.Barony = &许斯塔HustadBarony{}

func init() {
	f := BHustad许斯塔.(*许斯塔HustadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hustad",
		TitleName: "许斯塔",
		TitleCode: "b_hustad",
	}
}
