package steiermark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝汉贝格BehambergBarony struct {
	feud.BaseBarony
}

var BBehamberg贝汉贝格 feud.Barony = &贝汉贝格BehambergBarony{}

func init() {
	f := BBehamberg贝汉贝格.(*贝汉贝格BehambergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "behamberg",
		TitleName: "贝汉贝格",
		TitleCode: "b_behamberg",
	}
}
