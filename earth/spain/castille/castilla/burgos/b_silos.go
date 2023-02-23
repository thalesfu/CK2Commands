package burgos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡洛斯SilosBarony struct {
	feud.BaseBarony
}

var BSilos锡洛斯 feud.Barony = &锡洛斯SilosBarony{}

func init() {
	f := BSilos锡洛斯.(*锡洛斯SilosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "silos",
		TitleName: "锡洛斯",
		TitleCode: "b_silos",
	}
}
