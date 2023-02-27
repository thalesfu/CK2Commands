package trent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特伦托TrentoBarony struct {
	feud.BaseBarony
}

var BTrento特伦托 feud.Barony = &特伦托TrentoBarony{}

func init() {
    f := BTrento特伦托.(*特伦托TrentoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trento",
		TitleName: "特伦托",
		TitleCode: "b_trento",
	}
}
