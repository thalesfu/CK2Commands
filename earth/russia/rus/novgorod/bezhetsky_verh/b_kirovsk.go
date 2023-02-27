package bezhetsky_verh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基罗夫斯克KirovskBarony struct {
	feud.BaseBarony
}

var BKirovsk基罗夫斯克 feud.Barony = &基罗夫斯克KirovskBarony{}

func init() {
    f := BKirovsk基罗夫斯克.(*基罗夫斯克KirovskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirovsk",
		TitleName: "基罗夫斯克",
		TitleCode: "b_kirovsk",
	}
}
