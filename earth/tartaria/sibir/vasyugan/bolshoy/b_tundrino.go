package bolshoy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 通德里诺TundrinoBarony struct {
	feud.BaseBarony
}

var BTundrino通德里诺 feud.Barony = &通德里诺TundrinoBarony{}

func init() {
	f := BTundrino通德里诺.(*通德里诺TundrinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tundrino",
		TitleName: "通德里诺",
		TitleCode: "b_tundrino",
	}
}
