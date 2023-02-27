package ranthambore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢奥补罗SheopurBarony struct {
	feud.BaseBarony
}

var BSheopur谢奥补罗 feud.Barony = &谢奥补罗SheopurBarony{}

func init() {
    f := BSheopur谢奥补罗.(*谢奥补罗SheopurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sheopur",
		TitleName: "谢奥补罗",
		TitleCode: "b_sheopur",
	}
}
