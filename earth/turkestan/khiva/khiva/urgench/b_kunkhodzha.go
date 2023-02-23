package urgench

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昆霍贾KunkhodzhaBarony struct {
	feud.BaseBarony
}

var BKunkhodzha昆霍贾 feud.Barony = &昆霍贾KunkhodzhaBarony{}

func init() {
	f := BKunkhodzha昆霍贾.(*昆霍贾KunkhodzhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kunkhodzha",
		TitleName: "昆霍贾",
		TitleCode: "b_kunkhodzha",
	}
}
