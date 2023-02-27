package novogrodek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯洛尼姆SlonimBarony struct {
	feud.BaseBarony
}

var BSlonim斯洛尼姆 feud.Barony = &斯洛尼姆SlonimBarony{}

func init() {
    f := BSlonim斯洛尼姆.(*斯洛尼姆SlonimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "slonim",
		TitleName: "斯洛尼姆",
		TitleCode: "b_slonim",
	}
}
