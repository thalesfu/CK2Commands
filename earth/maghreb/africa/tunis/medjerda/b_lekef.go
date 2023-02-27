package medjerda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡夫LekefBarony struct {
	feud.BaseBarony
}

var BLekef卡夫 feud.Barony = &卡夫LekefBarony{}

func init() {
    f := BLekef卡夫.(*卡夫LekefBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lekef",
		TitleName: "卡夫",
		TitleCode: "b_lekef",
	}
}
