package arezzo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔托纳CortonaBarony struct {
	feud.BaseBarony
}

var BCortona科尔托纳 feud.Barony = &科尔托纳CortonaBarony{}

func init() {
    f := BCortona科尔托纳.(*科尔托纳CortonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cortona",
		TitleName: "科尔托纳",
		TitleCode: "b_cortona",
	}
}
