package gojjam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈贾姆GojjamBarony struct {
	feud.BaseBarony
}

var BGojjam戈贾姆 feud.Barony = &戈贾姆GojjamBarony{}

func init() {
    f := BGojjam戈贾姆.(*戈贾姆GojjamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gojjam",
		TitleName: "戈贾姆",
		TitleCode: "b_gojjam",
	}
}
