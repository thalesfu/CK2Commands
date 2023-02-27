package bauchi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丹马达洪DanmadahunBarony struct {
	feud.BaseBarony
}

var BDanmadahun丹马达洪 feud.Barony = &丹马达洪DanmadahunBarony{}

func init() {
    f := BDanmadahun丹马达洪.(*丹马达洪DanmadahunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "danmadahun",
		TitleName: "丹马达洪",
		TitleCode: "b_danmadahun",
	}
}
