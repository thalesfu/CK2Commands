package warzazat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔扎扎特WarzazatBarony struct {
	feud.BaseBarony
}

var BWarzazat瓦尔扎扎特 feud.Barony = &瓦尔扎扎特WarzazatBarony{}

func init() {
	f := BWarzazat瓦尔扎扎特.(*瓦尔扎扎特WarzazatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "warzazat",
		TitleName: "瓦尔扎扎特",
		TitleCode: "b_warzazat",
	}
}
