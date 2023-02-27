package skardu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀尔丘KharkooBarony struct {
	feud.BaseBarony
}

var BKharkoo喀尔丘 feud.Barony = &喀尔丘KharkooBarony{}

func init() {
    f := BKharkoo喀尔丘.(*喀尔丘KharkooBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kharkoo",
		TitleName: "喀尔丘",
		TitleCode: "b_kharkoo",
	}
}
