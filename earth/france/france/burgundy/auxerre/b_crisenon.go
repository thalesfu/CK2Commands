package auxerre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克里瑟农CrisenonBarony struct {
	feud.BaseBarony
}

var BCrisenon克里瑟农 feud.Barony = &克里瑟农CrisenonBarony{}

func init() {
	f := BCrisenon克里瑟农.(*克里瑟农CrisenonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "crisenon",
		TitleName: "克里瑟农",
		TitleCode: "b_crisenon",
	}
}
