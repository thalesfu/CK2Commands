package bamberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔姆贝格ColmbergBarony struct {
	feud.BaseBarony
}

var BColmberg科尔姆贝格 feud.Barony = &科尔姆贝格ColmbergBarony{}

func init() {
    f := BColmberg科尔姆贝格.(*科尔姆贝格ColmbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "colmberg",
		TitleName: "科尔姆贝格",
		TitleCode: "b_colmberg",
	}
}
