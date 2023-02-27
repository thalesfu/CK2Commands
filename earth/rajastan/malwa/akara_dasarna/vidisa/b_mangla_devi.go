package vidisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茫伽罗提毗Mangla_deviBarony struct {
	feud.BaseBarony
}

var BMangla_devi茫伽罗提毗 feud.Barony = &茫伽罗提毗Mangla_deviBarony{}

func init() {
    f := BMangla_devi茫伽罗提毗.(*茫伽罗提毗Mangla_deviBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mangla_devi",
		TitleName: "茫伽罗提毗",
		TitleCode: "b_mangla_devi",
	}
}
