package narke

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆索斯MosasBarony struct {
	feud.BaseBarony
}

var BMosas穆索斯 feud.Barony = &穆索斯MosasBarony{}

func init() {
	f := BMosas穆索斯.(*穆索斯MosasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mosas",
		TitleName: "穆索斯",
		TitleCode: "b_mosas",
	}
}
