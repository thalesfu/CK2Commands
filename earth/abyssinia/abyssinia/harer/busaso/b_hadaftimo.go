package busaso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈达费莫HadaftimoBarony struct {
	feud.BaseBarony
}

var BHadaftimo哈达费莫 feud.Barony = &哈达费莫HadaftimoBarony{}

func init() {
    f := BHadaftimo哈达费莫.(*哈达费莫HadaftimoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hadaftimo",
		TitleName: "哈达费莫",
		TitleCode: "b_hadaftimo",
	}
}
