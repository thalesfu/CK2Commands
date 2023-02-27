package chalons

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟尔SeurreBarony struct {
	feud.BaseBarony
}

var BSeurre瑟尔 feud.Barony = &瑟尔SeurreBarony{}

func init() {
    f := BSeurre瑟尔.(*瑟尔SeurreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "seurre",
		TitleName: "瑟尔",
		TitleCode: "b_seurre",
	}
}
