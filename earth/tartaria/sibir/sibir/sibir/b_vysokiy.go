package sibir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维索基VysokiyBarony struct {
	feud.BaseBarony
}

var BVysokiy维索基 feud.Barony = &维索基VysokiyBarony{}

func init() {
	f := BVysokiy维索基.(*维索基VysokiyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vysokiy",
		TitleName: "维索基",
		TitleCode: "b_vysokiy",
	}
}
