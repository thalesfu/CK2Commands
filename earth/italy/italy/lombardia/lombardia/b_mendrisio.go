package lombardia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 门德里西奥MendrisioBarony struct {
	feud.BaseBarony
}

var BMendrisio门德里西奥 feud.Barony = &门德里西奥MendrisioBarony{}

func init() {
    f := BMendrisio门德里西奥.(*门德里西奥MendrisioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mendrisio",
		TitleName: "门德里西奥",
		TitleCode: "b_mendrisio",
	}
}
