package lucca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔奇纳亚CalcinaiaBarony struct {
	feud.BaseBarony
}

var BCalcinaia卡尔奇纳亚 feud.Barony = &卡尔奇纳亚CalcinaiaBarony{}

func init() {
	f := BCalcinaia卡尔奇纳亚.(*卡尔奇纳亚CalcinaiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "calcinaia",
		TitleName: "卡尔奇纳亚",
		TitleCode: "b_calcinaia",
	}
}
