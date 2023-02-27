package krakowskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎科帕内ZakopaneBarony struct {
	feud.BaseBarony
}

var BZakopane扎科帕内 feud.Barony = &扎科帕内ZakopaneBarony{}

func init() {
    f := BZakopane扎科帕内.(*扎科帕内ZakopaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zakopane",
		TitleName: "扎科帕内",
		TitleCode: "b_zakopane",
	}
}
