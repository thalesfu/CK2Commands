package zatec

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特普利采TepliceBarony struct {
	feud.BaseBarony
}

var BTeplice特普利采 feud.Barony = &特普利采TepliceBarony{}

func init() {
    f := BTeplice特普利采.(*特普利采TepliceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "teplice",
		TitleName: "特普利采",
		TitleCode: "b_teplice",
	}
}
