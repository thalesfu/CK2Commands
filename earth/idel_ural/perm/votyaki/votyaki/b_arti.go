package votyaki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔季ArtiBarony struct {
	feud.BaseBarony
}

var BArti阿尔季 feud.Barony = &阿尔季ArtiBarony{}

func init() {
    f := BArti阿尔季.(*阿尔季ArtiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arti",
		TitleName: "阿尔季",
		TitleCode: "b_arti",
	}
}
