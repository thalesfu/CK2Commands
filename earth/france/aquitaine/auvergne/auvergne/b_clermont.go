package auvergne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克莱蒙ClermontBarony struct {
	feud.BaseBarony
}

var BClermont克莱蒙 feud.Barony = &克莱蒙ClermontBarony{}

func init() {
	f := BClermont克莱蒙.(*克莱蒙ClermontBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "clermont",
		TitleName: "克莱蒙",
		TitleCode: "b_clermont",
	}
}
