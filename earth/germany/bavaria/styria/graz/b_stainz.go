package graz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施泰因茨StainzBarony struct {
	feud.BaseBarony
}

var BStainz施泰因茨 feud.Barony = &施泰因茨StainzBarony{}

func init() {
	f := BStainz施泰因茨.(*施泰因茨StainzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stainz",
		TitleName: "施泰因茨",
		TitleCode: "b_stainz",
	}
}
