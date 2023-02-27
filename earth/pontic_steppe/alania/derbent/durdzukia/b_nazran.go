package durdzukia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳兹兰NazranBarony struct {
	feud.BaseBarony
}

var BNazran纳兹兰 feud.Barony = &纳兹兰NazranBarony{}

func init() {
    f := BNazran纳兹兰.(*纳兹兰NazranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nazran",
		TitleName: "纳兹兰",
		TitleCode: "b_nazran",
	}
}
