package taktse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡珠KazhuBarony struct {
	feud.BaseBarony
}

var BKazhu卡珠 feud.Barony = &卡珠KazhuBarony{}

func init() {
    f := BKazhu卡珠.(*卡珠KazhuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kazhu",
		TitleName: "卡珠",
		TitleCode: "b_kazhu",
	}
}
