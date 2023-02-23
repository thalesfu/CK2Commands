package lyubech

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷瑟耶LysyyeBarony struct {
	feud.BaseBarony
}

var BLysyye雷瑟耶 feud.Barony = &雷瑟耶LysyyeBarony{}

func init() {
	f := BLysyye雷瑟耶.(*雷瑟耶LysyyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lysyye",
		TitleName: "雷瑟耶",
		TitleCode: "b_lysyye",
	}
}
