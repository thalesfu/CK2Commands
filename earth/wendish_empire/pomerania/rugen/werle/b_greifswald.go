package werle

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格赖夫斯瓦尔德GreifswaldBarony struct {
	feud.BaseBarony
}

var BGreifswald格赖夫斯瓦尔德 feud.Barony = &格赖夫斯瓦尔德GreifswaldBarony{}

func init() {
    f := BGreifswald格赖夫斯瓦尔德.(*格赖夫斯瓦尔德GreifswaldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "greifswald",
		TitleName: "格赖夫斯瓦尔德",
		TitleCode: "b_greifswald",
	}
}
