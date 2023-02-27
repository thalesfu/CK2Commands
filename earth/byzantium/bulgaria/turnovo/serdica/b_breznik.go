package serdica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布雷兹尼克BreznikBarony struct {
	feud.BaseBarony
}

var BBreznik布雷兹尼克 feud.Barony = &布雷兹尼克BreznikBarony{}

func init() {
    f := BBreznik布雷兹尼克.(*布雷兹尼克BreznikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "breznik",
		TitleName: "布雷兹尼克",
		TitleCode: "b_breznik",
	}
}
