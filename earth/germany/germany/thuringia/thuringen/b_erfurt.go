package thuringen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔福特ErfurtBarony struct {
	feud.BaseBarony
}

var BErfurt埃尔福特 feud.Barony = &埃尔福特ErfurtBarony{}

func init() {
    f := BErfurt埃尔福特.(*埃尔福特ErfurtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erfurt",
		TitleName: "埃尔福特",
		TitleCode: "b_erfurt",
	}
}
