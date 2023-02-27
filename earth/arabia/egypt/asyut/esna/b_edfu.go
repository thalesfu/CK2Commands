package esna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃德富EdfuBarony struct {
	feud.BaseBarony
}

var BEdfu埃德富 feud.Barony = &埃德富EdfuBarony{}

func init() {
    f := BEdfu埃德富.(*埃德富EdfuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "edfu",
		TitleName: "埃德富",
		TitleCode: "b_edfu",
	}
}
