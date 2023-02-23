package mudgagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勿伽耆厘MudgagiriBarony struct {
	feud.BaseBarony
}

var BMudgagiri勿伽耆厘 feud.Barony = &勿伽耆厘MudgagiriBarony{}

func init() {
	f := BMudgagiri勿伽耆厘.(*勿伽耆厘MudgagiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mudgagiri",
		TitleName: "勿伽耆厘",
		TitleCode: "b_mudgagiri",
	}
}
