package upper_silesia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涅姆恰NiemczaBarony struct {
	feud.BaseBarony
}

var BNiemcza涅姆恰 feud.Barony = &涅姆恰NiemczaBarony{}

func init() {
    f := BNiemcza涅姆恰.(*涅姆恰NiemczaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "niemcza",
		TitleName: "涅姆恰",
		TitleCode: "b_niemcza",
	}
}
