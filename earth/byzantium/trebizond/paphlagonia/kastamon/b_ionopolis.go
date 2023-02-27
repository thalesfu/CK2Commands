package kastamon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊俄诺波利斯IonopolisBarony struct {
	feud.BaseBarony
}

var BIonopolis伊俄诺波利斯 feud.Barony = &伊俄诺波利斯IonopolisBarony{}

func init() {
    f := BIonopolis伊俄诺波利斯.(*伊俄诺波利斯IonopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ionopolis",
		TitleName: "伊俄诺波利斯",
		TitleCode: "b_ionopolis",
	}
}
