package zamora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费尔莫塞列FermoselleBarony struct {
	feud.BaseBarony
}

var BFermoselle费尔莫塞列 feud.Barony = &费尔莫塞列FermoselleBarony{}

func init() {
	f := BFermoselle费尔莫塞列.(*费尔莫塞列FermoselleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fermoselle",
		TitleName: "费尔莫塞列",
		TitleCode: "b_fermoselle",
	}
}
