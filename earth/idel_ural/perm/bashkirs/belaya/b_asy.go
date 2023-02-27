package belaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿瑟AsyBarony struct {
	feud.BaseBarony
}

var BAsy阿瑟 feud.Barony = &阿瑟AsyBarony{}

func init() {
    f := BAsy阿瑟.(*阿瑟AsyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asy",
		TitleName: "阿瑟",
		TitleCode: "b_asy",
	}
}
