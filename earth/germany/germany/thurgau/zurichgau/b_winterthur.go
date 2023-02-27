package zurichgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 温特图尔WinterthurBarony struct {
	feud.BaseBarony
}

var BWinterthur温特图尔 feud.Barony = &温特图尔WinterthurBarony{}

func init() {
    f := BWinterthur温特图尔.(*温特图尔WinterthurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "winterthur",
		TitleName: "温特图尔",
		TitleCode: "b_winterthur",
	}
}
