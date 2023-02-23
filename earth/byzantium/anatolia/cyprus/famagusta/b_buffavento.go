package famagusta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布法文托BuffaventoBarony struct {
	feud.BaseBarony
}

var BBuffavento布法文托 feud.Barony = &布法文托BuffaventoBarony{}

func init() {
	f := BBuffavento布法文托.(*布法文托BuffaventoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buffavento",
		TitleName: "布法文托",
		TitleCode: "b_buffavento",
	}
}
