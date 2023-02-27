package zahedan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡察鲁德KacharudBarony struct {
	feud.BaseBarony
}

var BKacharud卡察鲁德 feud.Barony = &卡察鲁德KacharudBarony{}

func init() {
    f := BKacharud卡察鲁德.(*卡察鲁德KacharudBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kacharud",
		TitleName: "卡察鲁德",
		TitleCode: "b_kacharud",
	}
}
