package abakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比斯卡姆扎BiskamzhaBarony struct {
	feud.BaseBarony
}

var BBiskamzha比斯卡姆扎 feud.Barony = &比斯卡姆扎BiskamzhaBarony{}

func init() {
    f := BBiskamzha比斯卡姆扎.(*比斯卡姆扎BiskamzhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "biskamzha",
		TitleName: "比斯卡姆扎",
		TitleCode: "b_biskamzha",
	}
}
