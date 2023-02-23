package kermanshah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎加瓦尔KangavarBarony struct {
	feud.BaseBarony
}

var BKangavar坎加瓦尔 feud.Barony = &坎加瓦尔KangavarBarony{}

func init() {
	f := BKangavar坎加瓦尔.(*坎加瓦尔KangavarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kangavar",
		TitleName: "坎加瓦尔",
		TitleCode: "b_kangavar",
	}
}
