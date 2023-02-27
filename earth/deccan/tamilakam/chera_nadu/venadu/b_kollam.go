package venadu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 故临KollamBarony struct {
	feud.BaseBarony
}

var BKollam故临 feud.Barony = &故临KollamBarony{}

func init() {
    f := BKollam故临.(*故临KollamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kollam",
		TitleName: "故临",
		TitleCode: "b_kollam",
	}
}
