package alcacer_do_sal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔瓦拉德AlvaladeBarony struct {
	feud.BaseBarony
}

var BAlvalade阿尔瓦拉德 feud.Barony = &阿尔瓦拉德AlvaladeBarony{}

func init() {
    f := BAlvalade阿尔瓦拉德.(*阿尔瓦拉德AlvaladeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alvalade",
		TitleName: "阿尔瓦拉德",
		TitleCode: "b_alvalade",
	}
}
