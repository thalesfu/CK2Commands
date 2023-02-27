package mangyul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 协格尔ShelkarBarony struct {
	feud.BaseBarony
}

var BShelkar协格尔 feud.Barony = &协格尔ShelkarBarony{}

func init() {
    f := BShelkar协格尔.(*协格尔ShelkarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shelkar",
		TitleName: "协格尔",
		TitleCode: "b_shelkar",
	}
}
