package jacwiez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 新格鲁多克NovogrudokBarony struct {
	feud.BaseBarony
}

var BNovogrudok新格鲁多克 feud.Barony = &新格鲁多克NovogrudokBarony{}

func init() {
    f := BNovogrudok新格鲁多克.(*新格鲁多克NovogrudokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "novogrudok",
		TitleName: "新格鲁多克",
		TitleCode: "b_novogrudok",
	}
}
