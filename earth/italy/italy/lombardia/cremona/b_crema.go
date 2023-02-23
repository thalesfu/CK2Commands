package cremona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克雷马CremaBarony struct {
	feud.BaseBarony
}

var BCrema克雷马 feud.Barony = &克雷马CremaBarony{}

func init() {
	f := BCrema克雷马.(*克雷马CremaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "crema",
		TitleName: "克雷马",
		TitleCode: "b_crema",
	}
}
