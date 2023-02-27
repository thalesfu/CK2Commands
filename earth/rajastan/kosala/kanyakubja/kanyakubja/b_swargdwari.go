package kanyakubja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莎罗伽堕利SwargdwariBarony struct {
	feud.BaseBarony
}

var BSwargdwari莎罗伽堕利 feud.Barony = &莎罗伽堕利SwargdwariBarony{}

func init() {
    f := BSwargdwari莎罗伽堕利.(*莎罗伽堕利SwargdwariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "swargdwari",
		TitleName: "莎罗伽堕利",
		TitleCode: "b_swargdwari",
	}
}
