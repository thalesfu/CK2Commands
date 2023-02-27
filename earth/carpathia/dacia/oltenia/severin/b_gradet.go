package severin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉德茨GradetBarony struct {
	feud.BaseBarony
}

var BGradet格拉德茨 feud.Barony = &格拉德茨GradetBarony{}

func init() {
    f := BGradet格拉德茨.(*格拉德茨GradetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gradet",
		TitleName: "格拉德茨",
		TitleCode: "b_gradet",
	}
}
