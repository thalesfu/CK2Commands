package viscaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣塞瓦斯蒂安SansebastianBarony struct {
	feud.BaseBarony
}

var BSansebastian圣塞瓦斯蒂安 feud.Barony = &圣塞瓦斯蒂安SansebastianBarony{}

func init() {
	f := BSansebastian圣塞瓦斯蒂安.(*圣塞瓦斯蒂安SansebastianBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sansebastian",
		TitleName: "圣塞瓦斯蒂安",
		TitleCode: "b_sansebastian",
	}
}
