package hy_many

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉纳达尔夫RathnadarveBarony struct {
	feud.BaseBarony
}

var BRathnadarve拉纳达尔夫 feud.Barony = &拉纳达尔夫RathnadarveBarony{}

func init() {
    f := BRathnadarve拉纳达尔夫.(*拉纳达尔夫RathnadarveBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rathnadarve",
		TitleName: "拉纳达尔夫",
		TitleCode: "b_rathnadarve",
	}
}
