package kartli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳里卡拉NarikalaBarony struct {
	feud.BaseBarony
}

var BNarikala纳里卡拉 feud.Barony = &纳里卡拉NarikalaBarony{}

func init() {
	f := BNarikala纳里卡拉.(*纳里卡拉NarikalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "narikala",
		TitleName: "纳里卡拉",
		TitleCode: "b_narikala",
	}
}
