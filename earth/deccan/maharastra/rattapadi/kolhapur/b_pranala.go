package kolhapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普罗纳拉PranalaBarony struct {
	feud.BaseBarony
}

var BPranala普罗纳拉 feud.Barony = &普罗纳拉PranalaBarony{}

func init() {
	f := BPranala普罗纳拉.(*普罗纳拉PranalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pranala",
		TitleName: "普罗纳拉",
		TitleCode: "b_pranala",
	}
}
