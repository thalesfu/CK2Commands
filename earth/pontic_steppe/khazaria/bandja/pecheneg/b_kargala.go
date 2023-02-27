package pecheneg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔加拉KargalaBarony struct {
	feud.BaseBarony
}

var BKargala卡尔加拉 feud.Barony = &卡尔加拉KargalaBarony{}

func init() {
    f := BKargala卡尔加拉.(*卡尔加拉KargalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kargala",
		TitleName: "卡尔加拉",
		TitleCode: "b_kargala",
	}
}
