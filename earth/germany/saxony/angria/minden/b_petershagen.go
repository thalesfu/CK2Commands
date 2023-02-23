package minden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 彼得斯哈根PetershagenBarony struct {
	feud.BaseBarony
}

var BPetershagen彼得斯哈根 feud.Barony = &彼得斯哈根PetershagenBarony{}

func init() {
	f := BPetershagen彼得斯哈根.(*彼得斯哈根PetershagenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "petershagen",
		TitleName: "彼得斯哈根",
		TitleCode: "b_petershagen",
	}
}
