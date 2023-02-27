package hamburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布伦斯比特尔BrunsbuttelBarony struct {
	feud.BaseBarony
}

var BBrunsbuttel布伦斯比特尔 feud.Barony = &布伦斯比特尔BrunsbuttelBarony{}

func init() {
    f := BBrunsbuttel布伦斯比特尔.(*布伦斯比特尔BrunsbuttelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brunsbuttel",
		TitleName: "布伦斯比特尔",
		TitleCode: "b_brunsbuttel",
	}
}
