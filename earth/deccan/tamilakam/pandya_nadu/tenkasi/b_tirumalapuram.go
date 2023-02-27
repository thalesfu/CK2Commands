package tenkasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提卢摩罗补楞TirumalapuramBarony struct {
	feud.BaseBarony
}

var BTirumalapuram提卢摩罗补楞 feud.Barony = &提卢摩罗补楞TirumalapuramBarony{}

func init() {
    f := BTirumalapuram提卢摩罗补楞.(*提卢摩罗补楞TirumalapuramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tirumalapuram",
		TitleName: "提卢摩罗补楞",
		TitleCode: "b_tirumalapuram",
	}
}
