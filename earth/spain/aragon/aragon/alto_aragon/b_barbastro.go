package alto_aragon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴瓦斯特罗BarbastroBarony struct {
	feud.BaseBarony
}

var BBarbastro巴瓦斯特罗 feud.Barony = &巴瓦斯特罗BarbastroBarony{}

func init() {
    f := BBarbastro巴瓦斯特罗.(*巴瓦斯特罗BarbastroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barbastro",
		TitleName: "巴瓦斯特罗",
		TitleCode: "b_barbastro",
	}
}
