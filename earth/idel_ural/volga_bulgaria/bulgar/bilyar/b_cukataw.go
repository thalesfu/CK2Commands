package bilyar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茹凯陶CukatawBarony struct {
	feud.BaseBarony
}

var BCukataw茹凯陶 feud.Barony = &茹凯陶CukatawBarony{}

func init() {
    f := BCukataw茹凯陶.(*茹凯陶CukatawBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cukataw",
		TitleName: "茹凯陶",
		TitleCode: "b_cukataw",
	}
}
