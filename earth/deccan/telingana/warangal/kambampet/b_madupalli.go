package kambampet

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马杜帕利MadupalliBarony struct {
	feud.BaseBarony
}

var BMadupalli马杜帕利 feud.Barony = &马杜帕利MadupalliBarony{}

func init() {
    f := BMadupalli马杜帕利.(*马杜帕利MadupalliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "madupalli",
		TitleName: "马杜帕利",
		TitleCode: "b_madupalli",
	}
}
