package thouars

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥洛讷OlonneBarony struct {
	feud.BaseBarony
}

var BOlonne奥洛讷 feud.Barony = &奥洛讷OlonneBarony{}

func init() {
    f := BOlonne奥洛讷.(*奥洛讷OlonneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olonne",
		TitleName: "奥洛讷",
		TitleCode: "b_olonne",
	}
}
