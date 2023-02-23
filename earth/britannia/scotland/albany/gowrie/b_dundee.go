package gowrie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓迪DundeeBarony struct {
	feud.BaseBarony
}

var BDundee邓迪 feud.Barony = &邓迪DundeeBarony{}

func init() {
	f := BDundee邓迪.(*邓迪DundeeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dundee",
		TitleName: "邓迪",
		TitleCode: "b_dundee",
	}
}
