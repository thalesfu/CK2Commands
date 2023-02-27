package surrey

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨瑟克SouthwarkBarony struct {
	feud.BaseBarony
}

var BSouthwark萨瑟克 feud.Barony = &萨瑟克SouthwarkBarony{}

func init() {
    f := BSouthwark萨瑟克.(*萨瑟克SouthwarkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "southwark",
		TitleName: "萨瑟克",
		TitleCode: "b_southwark",
	}
}
