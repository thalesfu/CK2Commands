package gaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那烂陀NalandaBarony struct {
	feud.BaseBarony
}

var BNalanda那烂陀 feud.Barony = &那烂陀NalandaBarony{}

func init() {
	f := BNalanda那烂陀.(*那烂陀NalandaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nalanda",
		TitleName: "那烂陀",
		TitleCode: "b_nalanda",
	}
}
