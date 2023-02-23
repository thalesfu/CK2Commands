package lhunzhub

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 松盘SumchengBarony struct {
	feud.BaseBarony
}

var BSumcheng松盘 feud.Barony = &松盘SumchengBarony{}

func init() {
	f := BSumcheng松盘.(*松盘SumchengBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sumcheng",
		TitleName: "松盘",
		TitleCode: "b_sumcheng",
	}
}
