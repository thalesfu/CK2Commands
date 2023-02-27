package tortosa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰卜莱JabalaBarony struct {
	feud.BaseBarony
}

var BJabala杰卜莱 feud.Barony = &杰卜莱JabalaBarony{}

func init() {
    f := BJabala杰卜莱.(*杰卜莱JabalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jabala",
		TitleName: "杰卜莱",
		TitleCode: "b_jabala",
	}
}
