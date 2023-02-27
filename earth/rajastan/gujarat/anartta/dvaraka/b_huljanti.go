package dvaraka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 呼旬提HuljantiBarony struct {
	feud.BaseBarony
}

var BHuljanti呼旬提 feud.Barony = &呼旬提HuljantiBarony{}

func init() {
    f := BHuljanti呼旬提.(*呼旬提HuljantiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "huljanti",
		TitleName: "呼旬提",
		TitleCode: "b_huljanti",
	}
}
