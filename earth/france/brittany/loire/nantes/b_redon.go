package nantes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒东RedonBarony struct {
	feud.BaseBarony
}

var BRedon勒东 feud.Barony = &勒东RedonBarony{}

func init() {
    f := BRedon勒东.(*勒东RedonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "redon",
		TitleName: "勒东",
		TitleCode: "b_redon",
	}
}
