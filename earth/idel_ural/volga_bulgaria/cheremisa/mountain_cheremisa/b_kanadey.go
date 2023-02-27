package mountain_cheremisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡纳杰伊KanadeyBarony struct {
	feud.BaseBarony
}

var BKanadey卡纳杰伊 feud.Barony = &卡纳杰伊KanadeyBarony{}

func init() {
    f := BKanadey卡纳杰伊.(*卡纳杰伊KanadeyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kanadey",
		TitleName: "卡纳杰伊",
		TitleCode: "b_kanadey",
	}
}
