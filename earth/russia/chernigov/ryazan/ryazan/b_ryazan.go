package ryazan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梁赞RyazanBarony struct {
	feud.BaseBarony
}

var BRyazan梁赞 feud.Barony = &梁赞RyazanBarony{}

func init() {
    f := BRyazan梁赞.(*梁赞RyazanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ryazan",
		TitleName: "梁赞",
		TitleCode: "b_ryazan",
	}
}
