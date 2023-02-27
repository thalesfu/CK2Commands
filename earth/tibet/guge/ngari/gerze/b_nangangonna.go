package gerze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那咔沃那NangangonnaBarony struct {
	feud.BaseBarony
}

var BNangangonna那咔沃那 feud.Barony = &那咔沃那NangangonnaBarony{}

func init() {
    f := BNangangonna那咔沃那.(*那咔沃那NangangonnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nangangonna",
		TitleName: "那咔沃那",
		TitleCode: "b_nangangonna",
	}
}
