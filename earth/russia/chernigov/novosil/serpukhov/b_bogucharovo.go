package serpukhov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博古恰罗沃BogucharovoBarony struct {
	feud.BaseBarony
}

var BBogucharovo博古恰罗沃 feud.Barony = &博古恰罗沃BogucharovoBarony{}

func init() {
    f := BBogucharovo博古恰罗沃.(*博古恰罗沃BogucharovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bogucharovo",
		TitleName: "博古恰罗沃",
		TitleCode: "b_bogucharovo",
	}
}
