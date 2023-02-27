package siwistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古奢GujaBarony struct {
	feud.BaseBarony
}

var BGuja古奢 feud.Barony = &古奢GujaBarony{}

func init() {
    f := BGuja古奢.(*古奢GujaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guja",
		TitleName: "古奢",
		TitleCode: "b_guja",
	}
}
