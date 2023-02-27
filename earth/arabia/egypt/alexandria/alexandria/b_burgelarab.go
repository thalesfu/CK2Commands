package alexandria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉伯堡BurgelarabBarony struct {
	feud.BaseBarony
}

var BBurgelarab阿拉伯堡 feud.Barony = &阿拉伯堡BurgelarabBarony{}

func init() {
    f := BBurgelarab阿拉伯堡.(*阿拉伯堡BurgelarabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burgelarab",
		TitleName: "阿拉伯堡",
		TitleCode: "b_burgelarab",
	}
}
