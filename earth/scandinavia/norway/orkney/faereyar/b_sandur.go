package faereyar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑杜尔SandurBarony struct {
	feud.BaseBarony
}

var BSandur桑杜尔 feud.Barony = &桑杜尔SandurBarony{}

func init() {
    f := BSandur桑杜尔.(*桑杜尔SandurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sandur",
		TitleName: "桑杜尔",
		TitleCode: "b_sandur",
	}
}
