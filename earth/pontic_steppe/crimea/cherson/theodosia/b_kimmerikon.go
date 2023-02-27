package theodosia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基梅里孔KimmerikonBarony struct {
	feud.BaseBarony
}

var BKimmerikon基梅里孔 feud.Barony = &基梅里孔KimmerikonBarony{}

func init() {
    f := BKimmerikon基梅里孔.(*基梅里孔KimmerikonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kimmerikon",
		TitleName: "基梅里孔",
		TitleCode: "b_kimmerikon",
	}
}
