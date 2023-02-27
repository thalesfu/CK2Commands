package soli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马格拉伊MaglajBarony struct {
	feud.BaseBarony
}

var BMaglaj马格拉伊 feud.Barony = &马格拉伊MaglajBarony{}

func init() {
    f := BMaglaj马格拉伊.(*马格拉伊MaglajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maglaj",
		TitleName: "马格拉伊",
		TitleCode: "b_maglaj",
	}
}
