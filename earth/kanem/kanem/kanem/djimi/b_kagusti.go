package djimi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡古斯提KagustiBarony struct {
	feud.BaseBarony
}

var BKagusti卡古斯提 feud.Barony = &卡古斯提KagustiBarony{}

func init() {
	f := BKagusti卡古斯提.(*卡古斯提KagustiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kagusti",
		TitleName: "卡古斯提",
		TitleCode: "b_kagusti",
	}
}
