package chester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切斯特ChesterBarony struct {
	feud.BaseBarony
}

var BChester切斯特 feud.Barony = &切斯特ChesterBarony{}

func init() {
    f := BChester切斯特.(*切斯特ChesterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chester",
		TitleName: "切斯特",
		TitleCode: "b_chester",
	}
}
