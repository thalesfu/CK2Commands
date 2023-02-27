package mesembria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尔奇多尔ValchidolBarony struct {
	feud.BaseBarony
}

var BValchidol沃尔奇多尔 feud.Barony = &沃尔奇多尔ValchidolBarony{}

func init() {
    f := BValchidol沃尔奇多尔.(*沃尔奇多尔ValchidolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valchidol",
		TitleName: "沃尔奇多尔",
		TitleCode: "b_valchidol",
	}
}
