package lancaster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰开斯特LancasterBarony struct {
	feud.BaseBarony
}

var BLancaster兰开斯特 feud.Barony = &兰开斯特LancasterBarony{}

func init() {
    f := BLancaster兰开斯特.(*兰开斯特LancasterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lancaster",
		TitleName: "兰开斯特",
		TitleCode: "b_lancaster",
	}
}
