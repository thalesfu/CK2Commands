package chuy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比什凯克BishkekBarony struct {
	feud.BaseBarony
}

var BBishkek比什凯克 feud.Barony = &比什凯克BishkekBarony{}

func init() {
	f := BBishkek比什凯克.(*比什凯克BishkekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bishkek",
		TitleName: "比什凯克",
		TitleCode: "b_bishkek",
	}
}
