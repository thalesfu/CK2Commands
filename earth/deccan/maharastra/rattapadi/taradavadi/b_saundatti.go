package taradavadi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏犍陀伐底SaundattiBarony struct {
	feud.BaseBarony
}

var BSaundatti苏犍陀伐底 feud.Barony = &苏犍陀伐底SaundattiBarony{}

func init() {
	f := BSaundatti苏犍陀伐底.(*苏犍陀伐底SaundattiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saundatti",
		TitleName: "苏犍陀伐底",
		TitleCode: "b_saundatti",
	}
}
