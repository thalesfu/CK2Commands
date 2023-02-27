package birlad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰库奇TecuciBarony struct {
	feud.BaseBarony
}

var BTecuci泰库奇 feud.Barony = &泰库奇TecuciBarony{}

func init() {
    f := BTecuci泰库奇.(*泰库奇TecuciBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tecuci",
		TitleName: "泰库奇",
		TitleCode: "b_tecuci",
	}
}
