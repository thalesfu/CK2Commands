package dhofar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈赞HarzanBarony struct {
	feud.BaseBarony
}

var BHarzan哈赞 feud.Barony = &哈赞HarzanBarony{}

func init() {
    f := BHarzan哈赞.(*哈赞HarzanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harzan",
		TitleName: "哈赞",
		TitleCode: "b_harzan",
	}
}
