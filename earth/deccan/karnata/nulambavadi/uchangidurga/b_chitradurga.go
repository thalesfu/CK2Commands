package uchangidurga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 质多罗突伽ChitradurgaBarony struct {
	feud.BaseBarony
}

var BChitradurga质多罗突伽 feud.Barony = &质多罗突伽ChitradurgaBarony{}

func init() {
    f := BChitradurga质多罗突伽.(*质多罗突伽ChitradurgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chitradurga",
		TitleName: "质多罗突伽",
		TitleCode: "b_chitradurga",
	}
}
