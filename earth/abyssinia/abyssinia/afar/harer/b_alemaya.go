package harer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿勒马亚AlemayaBarony struct {
	feud.BaseBarony
}

var BAlemaya阿勒马亚 feud.Barony = &阿勒马亚AlemayaBarony{}

func init() {
	f := BAlemaya阿勒马亚.(*阿勒马亚AlemayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alemaya",
		TitleName: "阿勒马亚",
		TitleCode: "b_alemaya",
	}
}
