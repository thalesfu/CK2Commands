package tripoli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜特龙BoutronBarony struct {
	feud.BaseBarony
}

var BBoutron拜特龙 feud.Barony = &拜特龙BoutronBarony{}

func init() {
	f := BBoutron拜特龙.(*拜特龙BoutronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boutron",
		TitleName: "拜特龙",
		TitleCode: "b_boutron",
	}
}
