package ani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米贾那贝德MidjnaberdBarony struct {
	feud.BaseBarony
}

var BMidjnaberd米贾那贝德 feud.Barony = &米贾那贝德MidjnaberdBarony{}

func init() {
    f := BMidjnaberd米贾那贝德.(*米贾那贝德MidjnaberdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "midjnaberd",
		TitleName: "米贾那贝德",
		TitleCode: "b_midjnaberd",
	}
}
