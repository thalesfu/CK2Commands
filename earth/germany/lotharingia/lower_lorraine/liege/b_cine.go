package liege

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西内CineBarony struct {
	feud.BaseBarony
}

var BCine西内 feud.Barony = &西内CineBarony{}

func init() {
    f := BCine西内.(*西内CineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cine",
		TitleName: "西内",
		TitleCode: "b_cine",
	}
}
