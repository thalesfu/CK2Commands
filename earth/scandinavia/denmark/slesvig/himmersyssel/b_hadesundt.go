package himmersyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海松HadesundtBarony struct {
	feud.BaseBarony
}

var BHadesundt海松 feud.Barony = &海松HadesundtBarony{}

func init() {
	f := BHadesundt海松.(*海松HadesundtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hadesundt",
		TitleName: "海松",
		TitleCode: "b_hadesundt",
	}
}
