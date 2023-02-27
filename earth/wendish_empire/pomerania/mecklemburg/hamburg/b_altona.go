package hamburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔托纳AltonaBarony struct {
	feud.BaseBarony
}

var BAltona阿尔托纳 feud.Barony = &阿尔托纳AltonaBarony{}

func init() {
    f := BAltona阿尔托纳.(*阿尔托纳AltonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "altona",
		TitleName: "阿尔托纳",
		TitleCode: "b_altona",
	}
}
