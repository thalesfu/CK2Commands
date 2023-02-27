package selija

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔维艾VarviaiBarony struct {
	feud.BaseBarony
}

var BVarviai瓦尔维艾 feud.Barony = &瓦尔维艾VarviaiBarony{}

func init() {
    f := BVarviai瓦尔维艾.(*瓦尔维艾VarviaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "varviai",
		TitleName: "瓦尔维艾",
		TitleCode: "b_varviai",
	}
}
