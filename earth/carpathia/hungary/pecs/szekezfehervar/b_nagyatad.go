package szekezfehervar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑙吉奥塔德NagyatadBarony struct {
	feud.BaseBarony
}

var BNagyatad瑙吉奥塔德 feud.Barony = &瑙吉奥塔德NagyatadBarony{}

func init() {
	f := BNagyatad瑙吉奥塔德.(*瑙吉奥塔德NagyatadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagyatad",
		TitleName: "瑙吉奥塔德",
		TitleCode: "b_nagyatad",
	}
}
