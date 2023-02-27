package campulung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布兰BranBarony struct {
	feud.BaseBarony
}

var BBran布兰 feud.Barony = &布兰BranBarony{}

func init() {
    f := BBran布兰.(*布兰BranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bran",
		TitleName: "布兰",
		TitleCode: "b_bran",
	}
}
