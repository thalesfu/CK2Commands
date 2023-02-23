package nandana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯沃拉KhewraBarony struct {
	feud.BaseBarony
}

var BKhewra凯沃拉 feud.Barony = &凯沃拉KhewraBarony{}

func init() {
	f := BKhewra凯沃拉.(*凯沃拉KhewraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khewra",
		TitleName: "凯沃拉",
		TitleCode: "b_khewra",
	}
}
