package hijaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨穆尔SamurBarony struct {
	feud.BaseBarony
}

var BSamur萨穆尔 feud.Barony = &萨穆尔SamurBarony{}

func init() {
    f := BSamur萨穆尔.(*萨穆尔SamurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samur",
		TitleName: "萨穆尔",
		TitleCode: "b_samur",
	}
}
