package badajoz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓜雷尼亚GuarenaBarony struct {
	feud.BaseBarony
}

var BGuarena瓜雷尼亚 feud.Barony = &瓜雷尼亚GuarenaBarony{}

func init() {
	f := BGuarena瓜雷尼亚.(*瓜雷尼亚GuarenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guarena",
		TitleName: "瓜雷尼亚",
		TitleCode: "b_guarena",
	}
}
