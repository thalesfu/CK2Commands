package penthievre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩朗PeranBarony struct {
	feud.BaseBarony
}

var BPeran佩朗 feud.Barony = &佩朗PeranBarony{}

func init() {
	f := BPeran佩朗.(*佩朗PeranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "peran",
		TitleName: "佩朗",
		TitleCode: "b_peran",
	}
}
