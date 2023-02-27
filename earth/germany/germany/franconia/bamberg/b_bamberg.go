package bamberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班贝格BambergBarony struct {
	feud.BaseBarony
}

var BBamberg班贝格 feud.Barony = &班贝格BambergBarony{}

func init() {
    f := BBamberg班贝格.(*班贝格BambergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bamberg",
		TitleName: "班贝格",
		TitleCode: "b_bamberg",
	}
}
