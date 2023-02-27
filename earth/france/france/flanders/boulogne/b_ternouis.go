package boulogne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰努瓦斯TernouisBarony struct {
	feud.BaseBarony
}

var BTernouis泰努瓦斯 feud.Barony = &泰努瓦斯TernouisBarony{}

func init() {
    f := BTernouis泰努瓦斯.(*泰努瓦斯TernouisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ternouis",
		TitleName: "泰努瓦斯",
		TitleCode: "b_ternouis",
	}
}
