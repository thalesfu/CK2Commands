package graz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗因ReunBarony struct {
	feud.BaseBarony
}

var BReun罗因 feud.Barony = &罗因ReunBarony{}

func init() {
	f := BReun罗因.(*罗因ReunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reun",
		TitleName: "罗因",
		TitleCode: "b_reun",
	}
}
