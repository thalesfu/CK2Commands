package yungguan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 寿昌ShouchangBarony struct {
	feud.BaseBarony
}

var BShouchang寿昌 feud.Barony = &寿昌ShouchangBarony{}

func init() {
	f := BShouchang寿昌.(*寿昌ShouchangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shouchang",
		TitleName: "寿昌",
		TitleCode: "b_shouchang",
	}
}
