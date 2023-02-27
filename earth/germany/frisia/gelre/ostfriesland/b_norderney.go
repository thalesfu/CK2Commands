package ostfriesland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺德奈NorderneyBarony struct {
	feud.BaseBarony
}

var BNorderney诺德奈 feud.Barony = &诺德奈NorderneyBarony{}

func init() {
    f := BNorderney诺德奈.(*诺德奈NorderneyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "norderney",
		TitleName: "诺德奈",
		TitleCode: "b_norderney",
	}
}
