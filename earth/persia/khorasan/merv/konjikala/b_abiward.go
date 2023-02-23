package konjikala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿比瓦尔德AbiwardBarony struct {
	feud.BaseBarony
}

var BAbiward阿比瓦尔德 feud.Barony = &阿比瓦尔德AbiwardBarony{}

func init() {
	f := BAbiward阿比瓦尔德.(*阿比瓦尔德AbiwardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abiward",
		TitleName: "阿比瓦尔德",
		TitleCode: "b_abiward",
	}
}
