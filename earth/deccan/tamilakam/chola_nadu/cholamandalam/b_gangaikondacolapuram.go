package cholamandalam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恒伽军荼朱罗补蓝GangaikondacolapuramBarony struct {
	feud.BaseBarony
}

var BGangaikondacolapuram恒伽军荼朱罗补蓝 feud.Barony = &恒伽军荼朱罗补蓝GangaikondacolapuramBarony{}

func init() {
    f := BGangaikondacolapuram恒伽军荼朱罗补蓝.(*恒伽军荼朱罗补蓝GangaikondacolapuramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gangaikondacolapuram",
		TitleName: "恒伽军荼朱罗补蓝",
		TitleCode: "b_gangaikondacolapuram",
	}
}
