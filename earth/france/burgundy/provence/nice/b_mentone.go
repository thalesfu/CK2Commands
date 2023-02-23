package nice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 门托内MentoneBarony struct {
	feud.BaseBarony
}

var BMentone门托内 feud.Barony = &门托内MentoneBarony{}

func init() {
	f := BMentone门托内.(*门托内MentoneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mentone",
		TitleName: "门托内",
		TitleCode: "b_mentone",
	}
}
