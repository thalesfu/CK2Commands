package perche

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 隆尼LongnyBarony struct {
	feud.BaseBarony
}

var BLongny隆尼 feud.Barony = &隆尼LongnyBarony{}

func init() {
    f := BLongny隆尼.(*隆尼LongnyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "longny",
		TitleName: "隆尼",
		TitleCode: "b_longny",
	}
}
