package mzab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖尔达耶GhardaiaBarony struct {
	feud.BaseBarony
}

var BGhardaia盖尔达耶 feud.Barony = &盖尔达耶GhardaiaBarony{}

func init() {
    f := BGhardaia盖尔达耶.(*盖尔达耶GhardaiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ghardaia",
		TitleName: "盖尔达耶",
		TitleCode: "b_ghardaia",
	}
}
