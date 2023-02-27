package idatarainadu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勿伽罗MudgalBarony struct {
	feud.BaseBarony
}

var BMudgal勿伽罗 feud.Barony = &勿伽罗MudgalBarony{}

func init() {
    f := BMudgal勿伽罗.(*勿伽罗MudgalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mudgal",
		TitleName: "勿伽罗",
		TitleCode: "b_mudgal",
	}
}
