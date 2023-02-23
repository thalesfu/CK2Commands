package boulogne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿德尔ArdresBarony struct {
	feud.BaseBarony
}

var BArdres阿德尔 feud.Barony = &阿德尔ArdresBarony{}

func init() {
	f := BArdres阿德尔.(*阿德尔ArdresBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ardres",
		TitleName: "阿德尔",
		TitleCode: "b_ardres",
	}
}
