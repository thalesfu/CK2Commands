package northampton

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克罗兰CrowlandBarony struct {
	feud.BaseBarony
}

var BCrowland克罗兰 feud.Barony = &克罗兰CrowlandBarony{}

func init() {
    f := BCrowland克罗兰.(*克罗兰CrowlandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "crowland",
		TitleName: "克罗兰",
		TitleCode: "b_crowland",
	}
}
