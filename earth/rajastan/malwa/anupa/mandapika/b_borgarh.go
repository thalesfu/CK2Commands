package mandapika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 浦迦尔BorgarhBarony struct {
	feud.BaseBarony
}

var BBorgarh浦迦尔 feud.Barony = &浦迦尔BorgarhBarony{}

func init() {
	f := BBorgarh浦迦尔.(*浦迦尔BorgarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borgarh",
		TitleName: "浦迦尔",
		TitleCode: "b_borgarh",
	}
}
