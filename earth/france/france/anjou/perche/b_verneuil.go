package perche

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦尔讷伊VerneuilBarony struct {
	feud.BaseBarony
}

var BVerneuil韦尔讷伊 feud.Barony = &韦尔讷伊VerneuilBarony{}

func init() {
    f := BVerneuil韦尔讷伊.(*韦尔讷伊VerneuilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "verneuil",
		TitleName: "韦尔讷伊",
		TitleCode: "b_verneuil",
	}
}
