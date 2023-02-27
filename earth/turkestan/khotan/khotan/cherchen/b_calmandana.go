package cherchen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 折摩驮那CalmandanaBarony struct {
	feud.BaseBarony
}

var BCalmandana折摩驮那 feud.Barony = &折摩驮那CalmandanaBarony{}

func init() {
    f := BCalmandana折摩驮那.(*折摩驮那CalmandanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "calmandana",
		TitleName: "折摩驮那",
		TitleCode: "b_calmandana",
	}
}
