package constantine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖尔马GuelmaBarony struct {
	feud.BaseBarony
}

var BGuelma盖尔马 feud.Barony = &盖尔马GuelmaBarony{}

func init() {
	f := BGuelma盖尔马.(*盖尔马GuelmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guelma",
		TitleName: "盖尔马",
		TitleCode: "b_guelma",
	}
}
