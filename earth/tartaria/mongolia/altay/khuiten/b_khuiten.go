package khuiten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 辉腾KhuitenBarony struct {
	feud.BaseBarony
}

var BKhuiten辉腾 feud.Barony = &辉腾KhuitenBarony{}

func init() {
	f := BKhuiten辉腾.(*辉腾KhuitenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khuiten",
		TitleName: "辉腾",
		TitleCode: "b_khuiten",
	}
}
