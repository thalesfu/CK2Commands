package kudalasangama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科潘KoppamBarony struct {
	feud.BaseBarony
}

var BKoppam科潘 feud.Barony = &科潘KoppamBarony{}

func init() {
    f := BKoppam科潘.(*科潘KoppamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koppam",
		TitleName: "科潘",
		TitleCode: "b_koppam",
	}
}
