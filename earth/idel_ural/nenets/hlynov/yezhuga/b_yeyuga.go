package yezhuga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶尤加YeyugaBarony struct {
	feud.BaseBarony
}

var BYeyuga叶尤加 feud.Barony = &叶尤加YeyugaBarony{}

func init() {
    f := BYeyuga叶尤加.(*叶尤加YeyugaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yeyuga",
		TitleName: "叶尤加",
		TitleCode: "b_yeyuga",
	}
}
