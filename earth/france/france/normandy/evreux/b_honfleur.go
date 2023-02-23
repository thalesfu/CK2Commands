package evreux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 翁弗勒尔HonfleurBarony struct {
	feud.BaseBarony
}

var BHonfleur翁弗勒尔 feud.Barony = &翁弗勒尔HonfleurBarony{}

func init() {
	f := BHonfleur翁弗勒尔.(*翁弗勒尔HonfleurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "honfleur",
		TitleName: "翁弗勒尔",
		TitleCode: "b_honfleur",
	}
}
