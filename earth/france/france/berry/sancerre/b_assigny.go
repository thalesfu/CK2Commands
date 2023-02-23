package sancerre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿西尼AssignyBarony struct {
	feud.BaseBarony
}

var BAssigny阿西尼 feud.Barony = &阿西尼AssignyBarony{}

func init() {
	f := BAssigny阿西尼.(*阿西尼AssignyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "assigny",
		TitleName: "阿西尼",
		TitleCode: "b_assigny",
	}
}
