package loon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马斯特里赫特MaastrichtBarony struct {
	feud.BaseBarony
}

var BMaastricht马斯特里赫特 feud.Barony = &马斯特里赫特MaastrichtBarony{}

func init() {
    f := BMaastricht马斯特里赫特.(*马斯特里赫特MaastrichtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maastricht",
		TitleName: "马斯特里赫特",
		TitleCode: "b_maastricht",
	}
}
