package reval

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 座堂山ToompeaBarony struct {
	feud.BaseBarony
}

var BToompea座堂山 feud.Barony = &座堂山ToompeaBarony{}

func init() {
	f := BToompea座堂山.(*座堂山ToompeaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "toompea",
		TitleName: "座堂山",
		TitleCode: "b_toompea",
	}
}
