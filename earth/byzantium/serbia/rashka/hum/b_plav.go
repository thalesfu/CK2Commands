package hum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普拉夫PlavBarony struct {
	feud.BaseBarony
}

var BPlav普拉夫 feud.Barony = &普拉夫PlavBarony{}

func init() {
	f := BPlav普拉夫.(*普拉夫PlavBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "plav",
		TitleName: "普拉夫",
		TitleCode: "b_plav",
	}
}
