package trent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布鲁内克BruneckBarony struct {
	feud.BaseBarony
}

var BBruneck布鲁内克 feud.Barony = &布鲁内克BruneckBarony{}

func init() {
	f := BBruneck布鲁内克.(*布鲁内克BruneckBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bruneck",
		TitleName: "布鲁内克",
		TitleCode: "b_bruneck",
	}
}
