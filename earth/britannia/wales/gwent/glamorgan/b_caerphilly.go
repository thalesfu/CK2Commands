package glamorgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡菲利CaerphillyBarony struct {
	feud.BaseBarony
}

var BCaerphilly卡菲利 feud.Barony = &卡菲利CaerphillyBarony{}

func init() {
	f := BCaerphilly卡菲利.(*卡菲利CaerphillyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caerphilly",
		TitleName: "卡菲利",
		TitleCode: "b_caerphilly",
	}
}
