package skardu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贺菩劳KhapluBarony struct {
	feud.BaseBarony
}

var BKhaplu贺菩劳 feud.Barony = &贺菩劳KhapluBarony{}

func init() {
	f := BKhaplu贺菩劳.(*贺菩劳KhapluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khaplu",
		TitleName: "贺菩劳",
		TitleCode: "b_khaplu",
	}
}
