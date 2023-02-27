package siuntio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡温蒂奥SiuntioBarony struct {
	feud.BaseBarony
}

var BSiuntio锡温蒂奥 feud.Barony = &锡温蒂奥SiuntioBarony{}

func init() {
    f := BSiuntio锡温蒂奥.(*锡温蒂奥SiuntioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siuntio",
		TitleName: "锡温蒂奥",
		TitleCode: "b_siuntio",
	}
}
