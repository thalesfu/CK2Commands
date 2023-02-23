package chester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 楠特威奇NantwichBarony struct {
	feud.BaseBarony
}

var BNantwich楠特威奇 feud.Barony = &楠特威奇NantwichBarony{}

func init() {
	f := BNantwich楠特威奇.(*楠特威奇NantwichBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nantwich",
		TitleName: "楠特威奇",
		TitleCode: "b_nantwich",
	}
}
