package zachlumia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯托拉茨StolacBarony struct {
	feud.BaseBarony
}

var BStolac斯托拉茨 feud.Barony = &斯托拉茨StolacBarony{}

func init() {
    f := BStolac斯托拉茨.(*斯托拉茨StolacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stolac",
		TitleName: "斯托拉茨",
		TitleCode: "b_stolac",
	}
}
