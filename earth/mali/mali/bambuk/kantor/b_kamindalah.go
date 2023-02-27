package kantor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡明达拉KamindalahBarony struct {
	feud.BaseBarony
}

var BKamindalah卡明达拉 feud.Barony = &卡明达拉KamindalahBarony{}

func init() {
    f := BKamindalah卡明达拉.(*卡明达拉KamindalahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamindalah",
		TitleName: "卡明达拉",
		TitleCode: "b_kamindalah",
	}
}
