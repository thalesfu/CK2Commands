package kantor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 芒迪奥MandioBarony struct {
	feud.BaseBarony
}

var BMandio芒迪奥 feud.Barony = &芒迪奥MandioBarony{}

func init() {
	f := BMandio芒迪奥.(*芒迪奥MandioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mandio",
		TitleName: "芒迪奥",
		TitleCode: "b_mandio",
	}
}
