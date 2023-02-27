package damascus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大马士革DamascusBarony struct {
	feud.BaseBarony
}

var BDamascus大马士革 feud.Barony = &大马士革DamascusBarony{}

func init() {
    f := BDamascus大马士革.(*大马士革DamascusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "damascus",
		TitleName: "大马士革",
		TitleCode: "b_damascus",
	}
}
