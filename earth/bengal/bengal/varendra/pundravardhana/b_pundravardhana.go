package pundravardhana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奔那伐弹那PundravardhanaBarony struct {
	feud.BaseBarony
}

var BPundravardhana奔那伐弹那 feud.Barony = &奔那伐弹那PundravardhanaBarony{}

func init() {
    f := BPundravardhana奔那伐弹那.(*奔那伐弹那PundravardhanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pundravardhana",
		TitleName: "奔那伐弹那",
		TitleCode: "b_pundravardhana",
	}
}
