package zemigalians

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅若特内MezotneBarony struct {
	feud.BaseBarony
}

var BMezotne梅若特内 feud.Barony = &梅若特内MezotneBarony{}

func init() {
    f := BMezotne梅若特内.(*梅若特内MezotneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mezotne",
		TitleName: "梅若特内",
		TitleCode: "b_mezotne",
	}
}
