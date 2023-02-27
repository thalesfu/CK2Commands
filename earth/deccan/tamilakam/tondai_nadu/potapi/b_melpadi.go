package potapi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩婆提MelpadiBarony struct {
	feud.BaseBarony
}

var BMelpadi摩婆提 feud.Barony = &摩婆提MelpadiBarony{}

func init() {
    f := BMelpadi摩婆提.(*摩婆提MelpadiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "melpadi",
		TitleName: "摩婆提",
		TitleCode: "b_melpadi",
	}
}
