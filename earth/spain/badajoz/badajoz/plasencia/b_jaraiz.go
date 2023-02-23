package plasencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈赖斯JaraizBarony struct {
	feud.BaseBarony
}

var BJaraiz哈赖斯 feud.Barony = &哈赖斯JaraizBarony{}

func init() {
	f := BJaraiz哈赖斯.(*哈赖斯JaraizBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jaraiz",
		TitleName: "哈赖斯",
		TitleCode: "b_jaraiz",
	}
}
