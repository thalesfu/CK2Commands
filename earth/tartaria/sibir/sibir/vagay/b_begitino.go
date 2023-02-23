package vagay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别吉季诺BegitinoBarony struct {
	feud.BaseBarony
}

var BBegitino别吉季诺 feud.Barony = &别吉季诺BegitinoBarony{}

func init() {
	f := BBegitino别吉季诺.(*别吉季诺BegitinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "begitino",
		TitleName: "别吉季诺",
		TitleCode: "b_begitino",
	}
}
