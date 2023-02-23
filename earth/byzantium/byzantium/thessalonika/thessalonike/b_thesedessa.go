package thessalonike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃泽萨ThesedessaBarony struct {
	feud.BaseBarony
}

var BThesedessa埃泽萨 feud.Barony = &埃泽萨ThesedessaBarony{}

func init() {
	f := BThesedessa埃泽萨.(*埃泽萨ThesedessaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thesedessa",
		TitleName: "埃泽萨",
		TitleCode: "b_thesedessa",
	}
}
