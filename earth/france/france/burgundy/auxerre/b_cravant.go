package auxerre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉旺CravantBarony struct {
	feud.BaseBarony
}

var BCravant克拉旺 feud.Barony = &克拉旺CravantBarony{}

func init() {
    f := BCravant克拉旺.(*克拉旺CravantBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cravant",
		TitleName: "克拉旺",
		TitleCode: "b_cravant",
	}
}
