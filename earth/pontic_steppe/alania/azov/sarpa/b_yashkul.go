package sarpa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚什库利YashkulBarony struct {
	feud.BaseBarony
}

var BYashkul亚什库利 feud.Barony = &亚什库利YashkulBarony{}

func init() {
    f := BYashkul亚什库利.(*亚什库利YashkulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yashkul",
		TitleName: "亚什库利",
		TitleCode: "b_yashkul",
	}
}
