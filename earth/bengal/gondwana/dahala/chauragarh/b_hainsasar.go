package chauragarh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 含娑萨HainsasarBarony struct {
	feud.BaseBarony
}

var BHainsasar含娑萨 feud.Barony = &含娑萨HainsasarBarony{}

func init() {
    f := BHainsasar含娑萨.(*含娑萨HainsasarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hainsasar",
		TitleName: "含娑萨",
		TitleCode: "b_hainsasar",
	}
}
