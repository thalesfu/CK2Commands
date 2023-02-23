package moskva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅季希MytishchiBarony struct {
	feud.BaseBarony
}

var BMytishchi梅季希 feud.Barony = &梅季希MytishchiBarony{}

func init() {
	f := BMytishchi梅季希.(*梅季希MytishchiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mytishchi",
		TitleName: "梅季希",
		TitleCode: "b_mytishchi",
	}
}
