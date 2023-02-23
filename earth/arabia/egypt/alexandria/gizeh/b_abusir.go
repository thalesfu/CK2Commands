package gizeh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布西尔AbusirBarony struct {
	feud.BaseBarony
}

var BAbusir阿布西尔 feud.Barony = &阿布西尔AbusirBarony{}

func init() {
	f := BAbusir阿布西尔.(*阿布西尔AbusirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abusir",
		TitleName: "阿布西尔",
		TitleCode: "b_abusir",
	}
}
