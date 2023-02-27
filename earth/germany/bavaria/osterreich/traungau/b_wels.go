package traungau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦尔斯WelsBarony struct {
	feud.BaseBarony
}

var BWels韦尔斯 feud.Barony = &韦尔斯WelsBarony{}

func init() {
    f := BWels韦尔斯.(*韦尔斯WelsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wels",
		TitleName: "韦尔斯",
		TitleCode: "b_wels",
	}
}
