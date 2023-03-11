package bryansk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏波涅沃SuponevoBarony struct {
	feud.BaseBarony
}

var BSuponevo苏波涅沃 feud.Barony = &苏波涅沃SuponevoBarony{}

func init() {
    f := BSuponevo苏波涅沃.(*苏波涅沃SuponevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suponevo",
		TitleName: "苏波涅沃",
		TitleCode: "b_suponevo",
	}
}
