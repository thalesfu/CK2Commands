package idatarainadu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阇罗突伽JaladurgaBarony struct {
	feud.BaseBarony
}

var BJaladurga阇罗突伽 feud.Barony = &阇罗突伽JaladurgaBarony{}

func init() {
    f := BJaladurga阇罗突伽.(*阇罗突伽JaladurgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jaladurga",
		TitleName: "阇罗突伽",
		TitleCode: "b_jaladurga",
	}
}
