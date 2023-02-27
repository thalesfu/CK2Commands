package naldurg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那罗突伽NaldurgBarony struct {
	feud.BaseBarony
}

var BNaldurg那罗突伽 feud.Barony = &那罗突伽NaldurgBarony{}

func init() {
    f := BNaldurg那罗突伽.(*那罗突伽NaldurgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naldurg",
		TitleName: "那罗突伽",
		TitleCode: "b_naldurg",
	}
}
