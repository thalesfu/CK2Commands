package padova

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基奥贾ChioggiaBarony struct {
	feud.BaseBarony
}

var BChioggia基奥贾 feud.Barony = &基奥贾ChioggiaBarony{}

func init() {
	f := BChioggia基奥贾.(*基奥贾ChioggiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chioggia",
		TitleName: "基奥贾",
		TitleCode: "b_chioggia",
	}
}
