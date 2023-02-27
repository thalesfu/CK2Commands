package eilat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基土拉KeturaBarony struct {
	feud.BaseBarony
}

var BKetura基土拉 feud.Barony = &基土拉KeturaBarony{}

func init() {
    f := BKetura基土拉.(*基土拉KeturaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ketura",
		TitleName: "基土拉",
		TitleCode: "b_ketura",
	}
}
