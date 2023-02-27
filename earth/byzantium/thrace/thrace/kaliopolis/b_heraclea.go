package kaliopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫拉克利亚HeracleaBarony struct {
	feud.BaseBarony
}

var BHeraclea赫拉克利亚 feud.Barony = &赫拉克利亚HeracleaBarony{}

func init() {
    f := BHeraclea赫拉克利亚.(*赫拉克利亚HeracleaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "heraclea",
		TitleName: "赫拉克利亚",
		TitleCode: "b_heraclea",
	}
}
