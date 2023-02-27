package kuopio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦亚萨洛VaajasaloBarony struct {
	feud.BaseBarony
}

var BVaajasalo瓦亚萨洛 feud.Barony = &瓦亚萨洛VaajasaloBarony{}

func init() {
    f := BVaajasalo瓦亚萨洛.(*瓦亚萨洛VaajasaloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vaajasalo",
		TitleName: "瓦亚萨洛",
		TitleCode: "b_vaajasalo",
	}
}
