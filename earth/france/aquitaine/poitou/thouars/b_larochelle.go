package thouars

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉罗谢尔LarochelleBarony struct {
	feud.BaseBarony
}

var BLarochelle拉罗谢尔 feud.Barony = &拉罗谢尔LarochelleBarony{}

func init() {
    f := BLarochelle拉罗谢尔.(*拉罗谢尔LarochelleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "larochelle",
		TitleName: "拉罗谢尔",
		TitleCode: "b_larochelle",
	}
}
