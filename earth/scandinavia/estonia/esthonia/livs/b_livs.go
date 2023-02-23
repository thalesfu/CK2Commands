package livs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利夫斯LivsBarony struct {
	feud.BaseBarony
}

var BLivs利夫斯 feud.Barony = &利夫斯LivsBarony{}

func init() {
	f := BLivs利夫斯.(*利夫斯LivsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "livs",
		TitleName: "利夫斯",
		TitleCode: "b_livs",
	}
}
