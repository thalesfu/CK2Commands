package komi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈尔普KharpBarony struct {
	feud.BaseBarony
}

var BKharp哈尔普 feud.Barony = &哈尔普KharpBarony{}

func init() {
    f := BKharp哈尔普.(*哈尔普KharpBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kharp",
		TitleName: "哈尔普",
		TitleCode: "b_kharp",
	}
}
