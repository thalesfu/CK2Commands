package lhatok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扩达KorraBarony struct {
	feud.BaseBarony
}

var BKorra扩达 feud.Barony = &扩达KorraBarony{}

func init() {
	f := BKorra扩达.(*扩达KorraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korra",
		TitleName: "扩达",
		TitleCode: "b_korra",
	}
}
