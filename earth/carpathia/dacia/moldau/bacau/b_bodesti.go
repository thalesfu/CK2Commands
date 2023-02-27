package bacau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博代什蒂BodestiBarony struct {
	feud.BaseBarony
}

var BBodesti博代什蒂 feud.Barony = &博代什蒂BodestiBarony{}

func init() {
    f := BBodesti博代什蒂.(*博代什蒂BodestiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bodesti",
		TitleName: "博代什蒂",
		TitleCode: "b_bodesti",
	}
}
