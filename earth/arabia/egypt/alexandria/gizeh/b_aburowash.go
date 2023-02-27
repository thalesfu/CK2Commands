package gizeh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布鲁韦斯AburowashBarony struct {
	feud.BaseBarony
}

var BAburowash阿布鲁韦斯 feud.Barony = &阿布鲁韦斯AburowashBarony{}

func init() {
    f := BAburowash阿布鲁韦斯.(*阿布鲁韦斯AburowashBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aburowash",
		TitleName: "阿布鲁韦斯",
		TitleCode: "b_aburowash",
	}
}
