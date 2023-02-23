package parnakheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菩罗荼那BuldhanaBarony struct {
	feud.BaseBarony
}

var BBuldhana菩罗荼那 feud.Barony = &菩罗荼那BuldhanaBarony{}

func init() {
	f := BBuldhana菩罗荼那.(*菩罗荼那BuldhanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buldhana",
		TitleName: "菩罗荼那",
		TitleCode: "b_buldhana",
	}
}
