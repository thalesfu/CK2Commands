package skara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克塞瓦拉AxevallaBarony struct {
	feud.BaseBarony
}

var BAxevalla阿克塞瓦拉 feud.Barony = &阿克塞瓦拉AxevallaBarony{}

func init() {
    f := BAxevalla阿克塞瓦拉.(*阿克塞瓦拉AxevallaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "axevalla",
		TitleName: "阿克塞瓦拉",
		TitleCode: "b_axevalla",
	}
}
