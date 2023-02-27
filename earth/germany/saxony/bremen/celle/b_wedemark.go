package celle

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦德马克WedemarkBarony struct {
	feud.BaseBarony
}

var BWedemark韦德马克 feud.Barony = &韦德马克WedemarkBarony{}

func init() {
    f := BWedemark韦德马克.(*韦德马克WedemarkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wedemark",
		TitleName: "韦德马克",
		TitleCode: "b_wedemark",
	}
}
