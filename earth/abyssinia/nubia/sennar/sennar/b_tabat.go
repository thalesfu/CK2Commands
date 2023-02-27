package sennar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔巴特TabatBarony struct {
	feud.BaseBarony
}

var BTabat塔巴特 feud.Barony = &塔巴特TabatBarony{}

func init() {
    f := BTabat塔巴特.(*塔巴特TabatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tabat",
		TitleName: "塔巴特",
		TitleCode: "b_tabat",
	}
}
