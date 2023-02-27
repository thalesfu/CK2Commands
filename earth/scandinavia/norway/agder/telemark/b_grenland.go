package telemark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格伦兰GrenlandBarony struct {
	feud.BaseBarony
}

var BGrenland格伦兰 feud.Barony = &格伦兰GrenlandBarony{}

func init() {
    f := BGrenland格伦兰.(*格伦兰GrenlandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grenland",
		TitleName: "格伦兰",
		TitleCode: "b_grenland",
	}
}
