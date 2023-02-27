package hayya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿里阿布AliabBarony struct {
	feud.BaseBarony
}

var BAliab阿里阿布 feud.Barony = &阿里阿布AliabBarony{}

func init() {
    f := BAliab阿里阿布.(*阿里阿布AliabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aliab",
		TitleName: "阿里阿布",
		TitleCode: "b_aliab",
	}
}
