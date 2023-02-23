package esfahan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿卜亚内赫AbyanehBarony struct {
	feud.BaseBarony
}

var BAbyaneh阿卜亚内赫 feud.Barony = &阿卜亚内赫AbyanehBarony{}

func init() {
	f := BAbyaneh阿卜亚内赫.(*阿卜亚内赫AbyanehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abyaneh",
		TitleName: "阿卜亚内赫",
		TitleCode: "b_abyaneh",
	}
}
