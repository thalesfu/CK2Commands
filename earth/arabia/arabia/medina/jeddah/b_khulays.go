package jeddah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡莱斯KhulaysBarony struct {
	feud.BaseBarony
}

var BKhulays胡莱斯 feud.Barony = &胡莱斯KhulaysBarony{}

func init() {
	f := BKhulays胡莱斯.(*胡莱斯KhulaysBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khulays",
		TitleName: "胡莱斯",
		TitleCode: "b_khulays",
	}
}
