package slesvig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 森讷堡SonderborgBarony struct {
	feud.BaseBarony
}

var BSonderborg森讷堡 feud.Barony = &森讷堡SonderborgBarony{}

func init() {
    f := BSonderborg森讷堡.(*森讷堡SonderborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sonderborg",
		TitleName: "森讷堡",
		TitleCode: "b_sonderborg",
	}
}
