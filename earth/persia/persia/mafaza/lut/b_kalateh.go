package lut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉泰KalatehBarony struct {
	feud.BaseBarony
}

var BKalateh卡拉泰 feud.Barony = &卡拉泰KalatehBarony{}

func init() {
    f := BKalateh卡拉泰.(*卡拉泰KalatehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalateh",
		TitleName: "卡拉泰",
		TitleCode: "b_kalateh",
	}
}
