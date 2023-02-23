package pereyaslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚戈京YahotynBarony struct {
	feud.BaseBarony
}

var BYahotyn亚戈京 feud.Barony = &亚戈京YahotynBarony{}

func init() {
	f := BYahotyn亚戈京.(*亚戈京YahotynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yahotyn",
		TitleName: "亚戈京",
		TitleCode: "b_yahotyn",
	}
}
