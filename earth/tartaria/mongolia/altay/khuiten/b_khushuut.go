package khuiten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫舍特KhushuutBarony struct {
	feud.BaseBarony
}

var BKhushuut赫舍特 feud.Barony = &赫舍特KhushuutBarony{}

func init() {
    f := BKhushuut赫舍特.(*赫舍特KhushuutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khushuut",
		TitleName: "赫舍特",
		TitleCode: "b_khushuut",
	}
}
