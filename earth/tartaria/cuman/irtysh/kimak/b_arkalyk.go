package kimak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔卡雷克ArkalykBarony struct {
	feud.BaseBarony
}

var BArkalyk阿尔卡雷克 feud.Barony = &阿尔卡雷克ArkalykBarony{}

func init() {
    f := BArkalyk阿尔卡雷克.(*阿尔卡雷克ArkalykBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arkalyk",
		TitleName: "阿尔卡雷克",
		TitleCode: "b_arkalyk",
	}
}
