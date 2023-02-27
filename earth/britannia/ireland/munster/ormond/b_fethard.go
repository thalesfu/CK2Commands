package ormond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费萨德FethardBarony struct {
	feud.BaseBarony
}

var BFethard费萨德 feud.Barony = &费萨德FethardBarony{}

func init() {
    f := BFethard费萨德.(*费萨德FethardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fethard",
		TitleName: "费萨德",
		TitleCode: "b_fethard",
	}
}
