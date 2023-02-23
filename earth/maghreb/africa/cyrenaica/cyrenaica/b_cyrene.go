package cyrenaica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昔兰尼CyreneBarony struct {
	feud.BaseBarony
}

var BCyrene昔兰尼 feud.Barony = &昔兰尼CyreneBarony{}

func init() {
	f := BCyrene昔兰尼.(*昔兰尼CyreneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cyrene",
		TitleName: "昔兰尼",
		TitleCode: "b_cyrene",
	}
}
