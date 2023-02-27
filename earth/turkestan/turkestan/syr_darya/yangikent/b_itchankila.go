package yangikent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊钱卡拉ItchankilaBarony struct {
	feud.BaseBarony
}

var BItchankila伊钱卡拉 feud.Barony = &伊钱卡拉ItchankilaBarony{}

func init() {
    f := BItchankila伊钱卡拉.(*伊钱卡拉ItchankilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "itchankila",
		TitleName: "伊钱卡拉",
		TitleCode: "b_itchankila",
	}
}
