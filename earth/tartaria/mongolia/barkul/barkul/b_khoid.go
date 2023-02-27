package barkul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 辉特KhoidBarony struct {
	feud.BaseBarony
}

var BKhoid辉特 feud.Barony = &辉特KhoidBarony{}

func init() {
    f := BKhoid辉特.(*辉特KhoidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khoid",
		TitleName: "辉特",
		TitleCode: "b_khoid",
	}
}
