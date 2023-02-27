package winchester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 南安普敦SouthamptonBarony struct {
	feud.BaseBarony
}

var BSouthampton南安普敦 feud.Barony = &南安普敦SouthamptonBarony{}

func init() {
    f := BSouthampton南安普敦.(*南安普敦SouthamptonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "southampton",
		TitleName: "南安普敦",
		TitleCode: "b_southampton",
	}
}
