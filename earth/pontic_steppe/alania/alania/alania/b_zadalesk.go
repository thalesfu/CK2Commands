package alania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎达列斯克ZadaleskBarony struct {
	feud.BaseBarony
}

var BZadalesk扎达列斯克 feud.Barony = &扎达列斯克ZadaleskBarony{}

func init() {
    f := BZadalesk扎达列斯克.(*扎达列斯克ZadaleskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zadalesk",
		TitleName: "扎达列斯克",
		TitleCode: "b_zadalesk",
	}
}
