package markakol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔内克OrnekBarony struct {
	feud.BaseBarony
}

var BOrnek奥尔内克 feud.Barony = &奥尔内克OrnekBarony{}

func init() {
    f := BOrnek奥尔内克.(*奥尔内克OrnekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ornek",
		TitleName: "奥尔内克",
		TitleCode: "b_ornek",
	}
}
