package vidisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑支SanchiBarony struct {
	feud.BaseBarony
}

var BSanchi桑支 feud.Barony = &桑支SanchiBarony{}

func init() {
    f := BSanchi桑支.(*桑支SanchiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanchi",
		TitleName: "桑支",
		TitleCode: "b_sanchi",
	}
}
