package buchan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伯丁AberdeenBarony struct {
	feud.BaseBarony
}

var BAberdeen阿伯丁 feud.Barony = &阿伯丁AberdeenBarony{}

func init() {
    f := BAberdeen阿伯丁.(*阿伯丁AberdeenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aberdeen",
		TitleName: "阿伯丁",
		TitleCode: "b_aberdeen",
	}
}
