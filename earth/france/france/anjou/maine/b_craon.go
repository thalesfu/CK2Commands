package maine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克朗CraonBarony struct {
	feud.BaseBarony
}

var BCraon克朗 feud.Barony = &克朗CraonBarony{}

func init() {
	f := BCraon克朗.(*克朗CraonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "craon",
		TitleName: "克朗",
		TitleCode: "b_craon",
	}
}
