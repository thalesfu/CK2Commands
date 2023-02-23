package kurmanchal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瞻波婆多ChampawatBarony struct {
	feud.BaseBarony
}

var BChampawat瞻波婆多 feud.Barony = &瞻波婆多ChampawatBarony{}

func init() {
	f := BChampawat瞻波婆多.(*瞻波婆多ChampawatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "champawat",
		TitleName: "瞻波婆多",
		TitleCode: "b_champawat",
	}
}
