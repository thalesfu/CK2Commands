package smyrna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希俄ChioBarony struct {
	feud.BaseBarony
}

var BChio希俄 feud.Barony = &希俄ChioBarony{}

func init() {
	f := BChio希俄.(*希俄ChioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chio",
		TitleName: "希俄",
		TitleCode: "b_chio",
	}
}
