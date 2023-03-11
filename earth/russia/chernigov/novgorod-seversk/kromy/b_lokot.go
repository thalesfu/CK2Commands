package kromy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛科季LokotBarony struct {
	feud.BaseBarony
}

var BLokot洛科季 feud.Barony = &洛科季LokotBarony{}

func init() {
    f := BLokot洛科季.(*洛科季LokotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lokot",
		TitleName: "洛科季",
		TitleCode: "b_lokot",
	}
}
