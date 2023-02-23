package vatsagulma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑多伽罗摩SatgaonBarony struct {
	feud.BaseBarony
}

var BSatgaon娑多伽罗摩 feud.Barony = &娑多伽罗摩SatgaonBarony{}

func init() {
	f := BSatgaon娑多伽罗摩.(*娑多伽罗摩SatgaonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "satgaon",
		TitleName: "娑多伽罗摩",
		TitleCode: "b_satgaon",
	}
}
