package livs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 派尔努ParnuBarony struct {
	feud.BaseBarony
}

var BParnu派尔努 feud.Barony = &派尔努ParnuBarony{}

func init() {
    f := BParnu派尔努.(*派尔努ParnuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "parnu",
		TitleName: "派尔努",
		TitleCode: "b_parnu",
	}
}
