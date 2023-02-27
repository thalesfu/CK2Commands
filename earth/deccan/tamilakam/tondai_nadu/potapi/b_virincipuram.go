package potapi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗利西补楞VirincipuramBarony struct {
	feud.BaseBarony
}

var BVirincipuram毗利西补楞 feud.Barony = &毗利西补楞VirincipuramBarony{}

func init() {
    f := BVirincipuram毗利西补楞.(*毗利西补楞VirincipuramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "virincipuram",
		TitleName: "毗利西补楞",
		TitleCode: "b_virincipuram",
	}
}
