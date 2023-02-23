package arques

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿夫勒尔HarfleurBarony struct {
	feud.BaseBarony
}

var BHarfleur阿夫勒尔 feud.Barony = &阿夫勒尔HarfleurBarony{}

func init() {
	f := BHarfleur阿夫勒尔.(*阿夫勒尔HarfleurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harfleur",
		TitleName: "阿夫勒尔",
		TitleCode: "b_harfleur",
	}
}
