package mohadavasaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旬卢JhalodBarony struct {
	feud.BaseBarony
}

var BJhalod旬卢 feud.Barony = &旬卢JhalodBarony{}

func init() {
    f := BJhalod旬卢.(*旬卢JhalodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jhalod",
		TitleName: "旬卢",
		TitleCode: "b_jhalod",
	}
}
