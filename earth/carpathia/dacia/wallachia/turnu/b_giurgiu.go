package turnu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 久尔久GiurgiuBarony struct {
	feud.BaseBarony
}

var BGiurgiu久尔久 feud.Barony = &久尔久GiurgiuBarony{}

func init() {
	f := BGiurgiu久尔久.(*久尔久GiurgiuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "giurgiu",
		TitleName: "久尔久",
		TitleCode: "b_giurgiu",
	}
}
