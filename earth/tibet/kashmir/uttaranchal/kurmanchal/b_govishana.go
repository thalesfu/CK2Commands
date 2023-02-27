package kurmanchal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瞿毗霜那GovishanaBarony struct {
	feud.BaseBarony
}

var BGovishana瞿毗霜那 feud.Barony = &瞿毗霜那GovishanaBarony{}

func init() {
    f := BGovishana瞿毗霜那.(*瞿毗霜那GovishanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "govishana",
		TitleName: "瞿毗霜那",
		TitleCode: "b_govishana",
	}
}
