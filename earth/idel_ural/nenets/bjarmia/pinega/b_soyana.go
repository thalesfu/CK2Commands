package pinega

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索亚纳SoyanaBarony struct {
	feud.BaseBarony
}

var BSoyana索亚纳 feud.Barony = &索亚纳SoyanaBarony{}

func init() {
    f := BSoyana索亚纳.(*索亚纳SoyanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soyana",
		TitleName: "索亚纳",
		TitleCode: "b_soyana",
	}
}
