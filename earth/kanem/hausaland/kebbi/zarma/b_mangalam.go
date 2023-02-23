package zarma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 芒加拉姆MangalamBarony struct {
	feud.BaseBarony
}

var BMangalam芒加拉姆 feud.Barony = &芒加拉姆MangalamBarony{}

func init() {
	f := BMangalam芒加拉姆.(*芒加拉姆MangalamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mangalam",
		TitleName: "芒加拉姆",
		TitleCode: "b_mangalam",
	}
}
