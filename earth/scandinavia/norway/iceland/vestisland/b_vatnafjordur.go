package vestisland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦特纳菲厄泽VatnafjordurBarony struct {
	feud.BaseBarony
}

var BVatnafjordur瓦特纳菲厄泽 feud.Barony = &瓦特纳菲厄泽VatnafjordurBarony{}

func init() {
	f := BVatnafjordur瓦特纳菲厄泽.(*瓦特纳菲厄泽VatnafjordurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vatnafjordur",
		TitleName: "瓦特纳菲厄泽",
		TitleCode: "b_vatnafjordur",
	}
}
