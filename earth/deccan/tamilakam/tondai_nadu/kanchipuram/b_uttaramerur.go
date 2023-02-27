package kanchipuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤多罗摩罗UttaramerurBarony struct {
	feud.BaseBarony
}

var BUttaramerur尤多罗摩罗 feud.Barony = &尤多罗摩罗UttaramerurBarony{}

func init() {
    f := BUttaramerur尤多罗摩罗.(*尤多罗摩罗UttaramerurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uttaramerur",
		TitleName: "尤多罗摩罗",
		TitleCode: "b_uttaramerur",
	}
}
