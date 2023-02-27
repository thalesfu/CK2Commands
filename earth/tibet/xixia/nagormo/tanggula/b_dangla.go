package tanggula

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 当拉DanglaBarony struct {
	feud.BaseBarony
}

var BDangla当拉 feud.Barony = &当拉DanglaBarony{}

func init() {
    f := BDangla当拉.(*当拉DanglaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dangla",
		TitleName: "当拉",
		TitleCode: "b_dangla",
	}
}
