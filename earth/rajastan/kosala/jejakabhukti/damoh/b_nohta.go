package damoh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奴诃多NohtaBarony struct {
	feud.BaseBarony
}

var BNohta奴诃多 feud.Barony = &奴诃多NohtaBarony{}

func init() {
	f := BNohta奴诃多.(*奴诃多NohtaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nohta",
		TitleName: "奴诃多",
		TitleCode: "b_nohta",
	}
}
