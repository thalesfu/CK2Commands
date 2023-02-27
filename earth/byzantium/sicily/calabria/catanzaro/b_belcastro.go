package catanzaro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尔卡斯特罗BelcastroBarony struct {
	feud.BaseBarony
}

var BBelcastro贝尔卡斯特罗 feud.Barony = &贝尔卡斯特罗BelcastroBarony{}

func init() {
    f := BBelcastro贝尔卡斯特罗.(*贝尔卡斯特罗BelcastroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belcastro",
		TitleName: "贝尔卡斯特罗",
		TitleCode: "b_belcastro",
	}
}
