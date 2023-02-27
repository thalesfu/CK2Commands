package paphlagonia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨夫兰博卢SafranboluBarony struct {
	feud.BaseBarony
}

var BSafranbolu萨夫兰博卢 feud.Barony = &萨夫兰博卢SafranboluBarony{}

func init() {
    f := BSafranbolu萨夫兰博卢.(*萨夫兰博卢SafranboluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "safranbolu",
		TitleName: "萨夫兰博卢",
		TitleCode: "b_safranbolu",
	}
}
