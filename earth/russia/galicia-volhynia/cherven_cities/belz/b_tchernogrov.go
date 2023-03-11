package belz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔诺格罗夫TchernogrovBarony struct {
	feud.BaseBarony
}

var BTchernogrov切尔诺格罗夫 feud.Barony = &切尔诺格罗夫TchernogrovBarony{}

func init() {
    f := BTchernogrov切尔诺格罗夫.(*切尔诺格罗夫TchernogrovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tchernogrov",
		TitleName: "切尔诺格罗夫",
		TitleCode: "b_tchernogrov",
	}
}
