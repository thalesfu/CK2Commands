package tripoli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 季比莱特GibeletBarony struct {
	feud.BaseBarony
}

var BGibelet季比莱特 feud.Barony = &季比莱特GibeletBarony{}

func init() {
    f := BGibelet季比莱特.(*季比莱特GibeletBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gibelet",
		TitleName: "季比莱特",
		TitleCode: "b_gibelet",
	}
}
