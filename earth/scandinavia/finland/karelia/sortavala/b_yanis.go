package sortavala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚尼斯YanisBarony struct {
	feud.BaseBarony
}

var BYanis亚尼斯 feud.Barony = &亚尼斯YanisBarony{}

func init() {
    f := BYanis亚尼斯.(*亚尼斯YanisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yanis",
		TitleName: "亚尼斯",
		TitleCode: "b_yanis",
	}
}
