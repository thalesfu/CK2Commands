package sopron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哲尔GyorBarony struct {
	feud.BaseBarony
}

var BGyor哲尔 feud.Barony = &哲尔GyorBarony{}

func init() {
    f := BGyor哲尔.(*哲尔GyorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gyor",
		TitleName: "哲尔",
		TitleCode: "b_gyor",
	}
}
