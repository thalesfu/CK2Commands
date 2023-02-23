package tebessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罕西拉KhenchelaBarony struct {
	feud.BaseBarony
}

var BKhenchela罕西拉 feud.Barony = &罕西拉KhenchelaBarony{}

func init() {
	f := BKhenchela罕西拉.(*罕西拉KhenchelaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khenchela",
		TitleName: "罕西拉",
		TitleCode: "b_khenchela",
	}
}
