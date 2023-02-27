package granada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦尔马HuelmaBarony struct {
	feud.BaseBarony
}

var BHuelma韦尔马 feud.Barony = &韦尔马HuelmaBarony{}

func init() {
    f := BHuelma韦尔马.(*韦尔马HuelmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "huelma",
		TitleName: "韦尔马",
		TitleCode: "b_huelma",
	}
}
