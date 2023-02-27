package ostfriesland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃姆登EmdenBarony struct {
	feud.BaseBarony
}

var BEmden埃姆登 feud.Barony = &埃姆登EmdenBarony{}

func init() {
    f := BEmden埃姆登.(*埃姆登EmdenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "emden",
		TitleName: "埃姆登",
		TitleCode: "b_emden",
	}
}
