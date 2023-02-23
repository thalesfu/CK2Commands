package goa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖里亚GheriaBarony struct {
	feud.BaseBarony
}

var BGheria盖里亚 feud.Barony = &盖里亚GheriaBarony{}

func init() {
	f := BGheria盖里亚.(*盖里亚GheriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gheria",
		TitleName: "盖里亚",
		TitleCode: "b_gheria",
	}
}
