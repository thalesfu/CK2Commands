package bytow

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕尔霍维ParchowyBarony struct {
	feud.BaseBarony
}

var BParchowy帕尔霍维 feud.Barony = &帕尔霍维ParchowyBarony{}

func init() {
    f := BParchowy帕尔霍维.(*帕尔霍维ParchowyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "parchowy",
		TitleName: "帕尔霍维",
		TitleCode: "b_parchowy",
	}
}
