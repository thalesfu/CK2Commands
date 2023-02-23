package amdo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 岗尼GangnyiBarony struct {
	feud.BaseBarony
}

var BGangnyi岗尼 feud.Barony = &岗尼GangnyiBarony{}

func init() {
	f := BGangnyi岗尼.(*岗尼GangnyiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gangnyi",
		TitleName: "岗尼",
		TitleCode: "b_gangnyi",
	}
}
