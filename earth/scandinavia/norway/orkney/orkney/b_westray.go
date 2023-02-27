package orkney

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦斯特雷WestrayBarony struct {
	feud.BaseBarony
}

var BWestray韦斯特雷 feud.Barony = &韦斯特雷WestrayBarony{}

func init() {
    f := BWestray韦斯特雷.(*韦斯特雷WestrayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "westray",
		TitleName: "韦斯特雷",
		TitleCode: "b_westray",
	}
}
