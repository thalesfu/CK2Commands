package claudiopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克劳狄奥波利斯ClaudiopolisBarony struct {
	feud.BaseBarony
}

var BClaudiopolis克劳狄奥波利斯 feud.Barony = &克劳狄奥波利斯ClaudiopolisBarony{}

func init() {
	f := BClaudiopolis克劳狄奥波利斯.(*克劳狄奥波利斯ClaudiopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "claudiopolis",
		TitleName: "克劳狄奥波利斯",
		TitleCode: "b_claudiopolis",
	}
}
