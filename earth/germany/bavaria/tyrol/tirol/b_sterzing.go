package tirol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施泰尔青SterzingBarony struct {
	feud.BaseBarony
}

var BSterzing施泰尔青 feud.Barony = &施泰尔青SterzingBarony{}

func init() {
    f := BSterzing施泰尔青.(*施泰尔青SterzingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sterzing",
		TitleName: "施泰尔青",
		TitleCode: "b_sterzing",
	}
}
