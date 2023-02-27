package zanjan_abhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿卜哈尔AbharBarony struct {
	feud.BaseBarony
}

var BAbhar阿卜哈尔 feud.Barony = &阿卜哈尔AbharBarony{}

func init() {
    f := BAbhar阿卜哈尔.(*阿卜哈尔AbharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abhar",
		TitleName: "阿卜哈尔",
		TitleCode: "b_abhar",
	}
}
