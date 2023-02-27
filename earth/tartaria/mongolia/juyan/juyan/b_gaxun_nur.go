package juyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 嘎顺淖尔Gaxun_nurBarony struct {
	feud.BaseBarony
}

var BGaxun_nur嘎顺淖尔 feud.Barony = &嘎顺淖尔Gaxun_nurBarony{}

func init() {
    f := BGaxun_nur嘎顺淖尔.(*嘎顺淖尔Gaxun_nurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gaxun_nur",
		TitleName: "嘎顺淖尔",
		TitleCode: "b_gaxun_nur",
	}
}
