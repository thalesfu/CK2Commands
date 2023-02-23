package mantua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维尔吉利奥VirgilioBarony struct {
	feud.BaseBarony
}

var BVirgilio维尔吉利奥 feud.Barony = &维尔吉利奥VirgilioBarony{}

func init() {
	f := BVirgilio维尔吉利奥.(*维尔吉利奥VirgilioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "virgilio",
		TitleName: "维尔吉利奥",
		TitleCode: "b_virgilio",
	}
}
