package slesvig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海泽比HedebyBarony struct {
	feud.BaseBarony
}

var BHedeby海泽比 feud.Barony = &海泽比HedebyBarony{}

func init() {
	f := BHedeby海泽比.(*海泽比HedebyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hedeby",
		TitleName: "海泽比",
		TitleCode: "b_hedeby",
	}
}
