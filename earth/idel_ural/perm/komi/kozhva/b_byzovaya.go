package kozhva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝佐瓦亚ByzovayaBarony struct {
	feud.BaseBarony
}

var BByzovaya贝佐瓦亚 feud.Barony = &贝佐瓦亚ByzovayaBarony{}

func init() {
    f := BByzovaya贝佐瓦亚.(*贝佐瓦亚ByzovayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "byzovaya",
		TitleName: "贝佐瓦亚",
		TitleCode: "b_byzovaya",
	}
}
