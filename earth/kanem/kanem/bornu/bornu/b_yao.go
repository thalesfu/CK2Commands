package bornu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚奥YaoBarony struct {
	feud.BaseBarony
}

var BYao亚奥 feud.Barony = &亚奥YaoBarony{}

func init() {
	f := BYao亚奥.(*亚奥YaoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yao",
		TitleName: "亚奥",
		TitleCode: "b_yao",
	}
}
