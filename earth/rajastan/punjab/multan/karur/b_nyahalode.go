package karur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 若诃卢提NyahalodeBarony struct {
	feud.BaseBarony
}

var BNyahalode若诃卢提 feud.Barony = &若诃卢提NyahalodeBarony{}

func init() {
    f := BNyahalode若诃卢提.(*若诃卢提NyahalodeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nyahalode",
		TitleName: "若诃卢提",
		TitleCode: "b_nyahalode",
	}
}
