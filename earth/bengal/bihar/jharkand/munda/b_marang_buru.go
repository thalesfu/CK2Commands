package munda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩烂迦补卢Marang_buruBarony struct {
	feud.BaseBarony
}

var BMarang_buru摩烂迦补卢 feud.Barony = &摩烂迦补卢Marang_buruBarony{}

func init() {
    f := BMarang_buru摩烂迦补卢.(*摩烂迦补卢Marang_buruBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marang_buru",
		TitleName: "摩烂迦补卢",
		TitleCode: "b_marang_buru",
	}
}
