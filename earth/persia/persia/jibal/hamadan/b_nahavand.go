package hamadan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳哈万德NahavandBarony struct {
	feud.BaseBarony
}

var BNahavand纳哈万德 feud.Barony = &纳哈万德NahavandBarony{}

func init() {
    f := BNahavand纳哈万德.(*纳哈万德NahavandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nahavand",
		TitleName: "纳哈万德",
		TitleCode: "b_nahavand",
	}
}
