package gizeh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布吉拉伯AbughorobBarony struct {
	feud.BaseBarony
}

var BAbughorob阿布吉拉伯 feud.Barony = &阿布吉拉伯AbughorobBarony{}

func init() {
	f := BAbughorob阿布吉拉伯.(*阿布吉拉伯AbughorobBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abughorob",
		TitleName: "阿布吉拉伯",
		TitleCode: "b_abughorob",
	}
}
