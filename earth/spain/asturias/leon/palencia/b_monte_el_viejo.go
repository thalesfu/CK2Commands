package palencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔别霍山Monte_el_viejoBarony struct {
	feud.BaseBarony
}

var BMonte_el_viejo埃尔别霍山 feud.Barony = &埃尔别霍山Monte_el_viejoBarony{}

func init() {
    f := BMonte_el_viejo埃尔别霍山.(*埃尔别霍山Monte_el_viejoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monte_el_viejo",
		TitleName: "埃尔别霍山",
		TitleCode: "b_monte_el_viejo",
	}
}
