package plzen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多马日利采DomazliceBarony struct {
	feud.BaseBarony
}

var BDomazlice多马日利采 feud.Barony = &多马日利采DomazliceBarony{}

func init() {
    f := BDomazlice多马日利采.(*多马日利采DomazliceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "domazlice",
		TitleName: "多马日利采",
		TitleCode: "b_domazlice",
	}
}
