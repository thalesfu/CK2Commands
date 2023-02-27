package medjerda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡勒亚奈SilianaBarony struct {
	feud.BaseBarony
}

var BSiliana锡勒亚奈 feud.Barony = &锡勒亚奈SilianaBarony{}

func init() {
    f := BSiliana锡勒亚奈.(*锡勒亚奈SilianaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siliana",
		TitleName: "锡勒亚奈",
		TitleCode: "b_siliana",
	}
}
