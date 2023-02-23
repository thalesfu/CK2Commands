package kuwait

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 劳扎塔因ArrawdataynBarony struct {
	feud.BaseBarony
}

var BArrawdatayn劳扎塔因 feud.Barony = &劳扎塔因ArrawdataynBarony{}

func init() {
	f := BArrawdatayn劳扎塔因.(*劳扎塔因ArrawdataynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arrawdatayn",
		TitleName: "劳扎塔因",
		TitleCode: "b_arrawdatayn",
	}
}
