package reims

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙隆ChalonsBarony struct {
	feud.BaseBarony
}

var BChalons沙隆 feud.Barony = &沙隆ChalonsBarony{}

func init() {
    f := BChalons沙隆.(*沙隆ChalonsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chalons",
		TitleName: "沙隆",
		TitleCode: "b_chalons",
	}
}
