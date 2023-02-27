package gwent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纽波特NewportBarony struct {
	feud.BaseBarony
}

var BNewport纽波特 feud.Barony = &纽波特NewportBarony{}

func init() {
    f := BNewport纽波特.(*纽波特NewportBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "newport",
		TitleName: "纽波特",
		TitleCode: "b_newport",
	}
}
