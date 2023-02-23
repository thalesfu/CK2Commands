package magnitaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古谢沃GusevoBarony struct {
	feud.BaseBarony
}

var BGusevo古谢沃 feud.Barony = &古谢沃GusevoBarony{}

func init() {
	f := BGusevo古谢沃.(*古谢沃GusevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gusevo",
		TitleName: "古谢沃",
		TitleCode: "b_gusevo",
	}
}
