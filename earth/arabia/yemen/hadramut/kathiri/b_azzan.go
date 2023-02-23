package kathiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿赞AzzanBarony struct {
	feud.BaseBarony
}

var BAzzan阿赞 feud.Barony = &阿赞AzzanBarony{}

func init() {
	f := BAzzan阿赞.(*阿赞AzzanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "azzan",
		TitleName: "阿赞",
		TitleCode: "b_azzan",
	}
}
