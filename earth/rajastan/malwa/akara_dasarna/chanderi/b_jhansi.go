package chanderi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 占西JhansiBarony struct {
	feud.BaseBarony
}

var BJhansi占西 feud.Barony = &占西JhansiBarony{}

func init() {
    f := BJhansi占西.(*占西JhansiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jhansi",
		TitleName: "占西",
		TitleCode: "b_jhansi",
	}
}
