package kaliopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希斯图斯SestusBarony struct {
	feud.BaseBarony
}

var BSestus希斯图斯 feud.Barony = &希斯图斯SestusBarony{}

func init() {
	f := BSestus希斯图斯.(*希斯图斯SestusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sestus",
		TitleName: "希斯图斯",
		TitleCode: "b_sestus",
	}
}
