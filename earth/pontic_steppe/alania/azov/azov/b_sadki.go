package azov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨德基SadkiBarony struct {
	feud.BaseBarony
}

var BSadki萨德基 feud.Barony = &萨德基SadkiBarony{}

func init() {
    f := BSadki萨德基.(*萨德基SadkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sadki",
		TitleName: "萨德基",
		TitleCode: "b_sadki",
	}
}
