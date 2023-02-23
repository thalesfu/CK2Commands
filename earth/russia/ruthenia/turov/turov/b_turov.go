package turov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图罗夫TurovBarony struct {
	feud.BaseBarony
}

var BTurov图罗夫 feud.Barony = &图罗夫TurovBarony{}

func init() {
	f := BTurov图罗夫.(*图罗夫TurovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turov",
		TitleName: "图罗夫",
		TitleCode: "b_turov",
	}
}
