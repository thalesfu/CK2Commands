package vannes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦讷VannesBarony struct {
	feud.BaseBarony
}

var BVannes瓦讷 feud.Barony = &瓦讷VannesBarony{}

func init() {
    f := BVannes瓦讷.(*瓦讷VannesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vannes",
		TitleName: "瓦讷",
		TitleCode: "b_vannes",
	}
}
