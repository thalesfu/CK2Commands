package lubeck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦斯勒WesloBarony struct {
	feud.BaseBarony
}

var BWeslo韦斯勒 feud.Barony = &韦斯勒WesloBarony{}

func init() {
    f := BWeslo韦斯勒.(*韦斯勒WesloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "weslo",
		TitleName: "韦斯勒",
		TitleCode: "b_weslo",
	}
}
