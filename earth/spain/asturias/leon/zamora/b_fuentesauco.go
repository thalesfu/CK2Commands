package zamora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丰特萨乌科FuentesaucoBarony struct {
	feud.BaseBarony
}

var BFuentesauco丰特萨乌科 feud.Barony = &丰特萨乌科FuentesaucoBarony{}

func init() {
	f := BFuentesauco丰特萨乌科.(*丰特萨乌科FuentesaucoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fuentesauco",
		TitleName: "丰特萨乌科",
		TitleCode: "b_fuentesauco",
	}
}
