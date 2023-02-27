package sibir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊别尔_保加尔IberbolgarBarony struct {
	feud.BaseBarony
}

var BIberbolgar伊别尔_保加尔 feud.Barony = &伊别尔_保加尔IberbolgarBarony{}

func init() {
    f := BIberbolgar伊别尔_保加尔.(*伊别尔_保加尔IberbolgarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iberbolgar",
		TitleName: "伊别尔_保加尔",
		TitleCode: "b_iberbolgar",
	}
}
