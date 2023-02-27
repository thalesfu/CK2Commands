package kumtag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 小锅峪XiaoguoyuBarony struct {
	feud.BaseBarony
}

var BXiaoguoyu小锅峪 feud.Barony = &小锅峪XiaoguoyuBarony{}

func init() {
    f := BXiaoguoyu小锅峪.(*小锅峪XiaoguoyuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xiaoguoyu",
		TitleName: "小锅峪",
		TitleCode: "b_xiaoguoyu",
	}
}
