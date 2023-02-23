package frankfurt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡岑埃尔恩博根KatzenelnbogenBarony struct {
	feud.BaseBarony
}

var BKatzenelnbogen卡岑埃尔恩博根 feud.Barony = &卡岑埃尔恩博根KatzenelnbogenBarony{}

func init() {
	f := BKatzenelnbogen卡岑埃尔恩博根.(*卡岑埃尔恩博根KatzenelnbogenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "katzenelnbogen",
		TitleName: "卡岑埃尔恩博根",
		TitleCode: "b_katzenelnbogen",
	}
}
