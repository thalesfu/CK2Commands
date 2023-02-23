package valabhi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡霍尔SihorBarony struct {
	feud.BaseBarony
}

var BSihor锡霍尔 feud.Barony = &锡霍尔SihorBarony{}

func init() {
	f := BSihor锡霍尔.(*锡霍尔SihorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sihor",
		TitleName: "锡霍尔",
		TitleCode: "b_sihor",
	}
}
