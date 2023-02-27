package ayodhya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那祇湿伐罗那他NageshwarnathBarony struct {
	feud.BaseBarony
}

var BNageshwarnath那祇湿伐罗那他 feud.Barony = &那祇湿伐罗那他NageshwarnathBarony{}

func init() {
    f := BNageshwarnath那祇湿伐罗那他.(*那祇湿伐罗那他NageshwarnathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nageshwarnath",
		TitleName: "那祇湿伐罗那他",
		TitleCode: "b_nageshwarnath",
	}
}
