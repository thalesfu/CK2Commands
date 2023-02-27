package izhma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡贾布斯克SizyabskBarony struct {
	feud.BaseBarony
}

var BSizyabsk锡贾布斯克 feud.Barony = &锡贾布斯克SizyabskBarony{}

func init() {
    f := BSizyabsk锡贾布斯克.(*锡贾布斯克SizyabskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sizyabsk",
		TitleName: "锡贾布斯克",
		TitleCode: "b_sizyabsk",
	}
}
