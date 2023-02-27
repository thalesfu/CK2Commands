package tabaristan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 察姆诺ChamnooBarony struct {
	feud.BaseBarony
}

var BChamnoo察姆诺 feud.Barony = &察姆诺ChamnooBarony{}

func init() {
    f := BChamnoo察姆诺.(*察姆诺ChamnooBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chamnoo",
		TitleName: "察姆诺",
		TitleCode: "b_chamnoo",
	}
}
