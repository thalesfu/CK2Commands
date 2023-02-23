package slutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈梅利GomelBarony struct {
	feud.BaseBarony
}

var BGomel戈梅利 feud.Barony = &戈梅利GomelBarony{}

func init() {
	f := BGomel戈梅利.(*戈梅利GomelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gomel",
		TitleName: "戈梅利",
		TitleCode: "b_gomel",
	}
}
