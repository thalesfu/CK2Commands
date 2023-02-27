package khetaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 堪婆多KhambhatBarony struct {
	feud.BaseBarony
}

var BKhambhat堪婆多 feud.Barony = &堪婆多KhambhatBarony{}

func init() {
    f := BKhambhat堪婆多.(*堪婆多KhambhatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khambhat",
		TitleName: "堪婆多",
		TitleCode: "b_khambhat",
	}
}
