package merv

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古兰利GulanlyBarony struct {
	feud.BaseBarony
}

var BGulanly古兰利 feud.Barony = &古兰利GulanlyBarony{}

func init() {
    f := BGulanly古兰利.(*古兰利GulanlyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gulanly",
		TitleName: "古兰利",
		TitleCode: "b_gulanly",
	}
}
