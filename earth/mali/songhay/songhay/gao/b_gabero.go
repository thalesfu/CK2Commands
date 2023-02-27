package gao

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加贝罗GaberoBarony struct {
	feud.BaseBarony
}

var BGabero加贝罗 feud.Barony = &加贝罗GaberoBarony{}

func init() {
    f := BGabero加贝罗.(*加贝罗GaberoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gabero",
		TitleName: "加贝罗",
		TitleCode: "b_gabero",
	}
}
