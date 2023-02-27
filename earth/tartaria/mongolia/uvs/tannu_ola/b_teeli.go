package tannu_ola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰利TeeliBarony struct {
	feud.BaseBarony
}

var BTeeli泰利 feud.Barony = &泰利TeeliBarony{}

func init() {
    f := BTeeli泰利.(*泰利TeeliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "teeli",
		TitleName: "泰利",
		TitleCode: "b_teeli",
	}
}
