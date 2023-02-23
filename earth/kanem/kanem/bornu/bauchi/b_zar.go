package bauchi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎尔ZarBarony struct {
	feud.BaseBarony
}

var BZar扎尔 feud.Barony = &扎尔ZarBarony{}

func init() {
	f := BZar扎尔.(*扎尔ZarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zar",
		TitleName: "扎尔",
		TitleCode: "b_zar",
	}
}
