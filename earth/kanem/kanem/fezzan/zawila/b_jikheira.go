package zawila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾哈拉JikheiraBarony struct {
	feud.BaseBarony
}

var BJikheira贾哈拉 feud.Barony = &贾哈拉JikheiraBarony{}

func init() {
	f := BJikheira贾哈拉.(*贾哈拉JikheiraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jikheira",
		TitleName: "贾哈拉",
		TitleCode: "b_jikheira",
	}
}
