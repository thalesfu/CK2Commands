package pfalz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特里费尔斯TrifelsBarony struct {
	feud.BaseBarony
}

var BTrifels特里费尔斯 feud.Barony = &特里费尔斯TrifelsBarony{}

func init() {
    f := BTrifels特里费尔斯.(*特里费尔斯TrifelsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trifels",
		TitleName: "特里费尔斯",
		TitleCode: "b_trifels",
	}
}
