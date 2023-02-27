package tribandapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耆那多城JaintpuriBarony struct {
	feud.BaseBarony
}

var BJaintpuri耆那多城 feud.Barony = &耆那多城JaintpuriBarony{}

func init() {
    f := BJaintpuri耆那多城.(*耆那多城JaintpuriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jaintpuri",
		TitleName: "耆那多城",
		TitleCode: "b_jaintpuri",
	}
}
