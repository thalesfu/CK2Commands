package burkhan_buudai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 车勒TseelBarony struct {
	feud.BaseBarony
}

var BTseel车勒 feud.Barony = &车勒TseelBarony{}

func init() {
    f := BTseel车勒.(*车勒TseelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tseel",
		TitleName: "车勒",
		TitleCode: "b_tseel",
	}
}
