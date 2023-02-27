package nalut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富尔塞塔FursataBarony struct {
	feud.BaseBarony
}

var BFursata富尔塞塔 feud.Barony = &富尔塞塔FursataBarony{}

func init() {
    f := BFursata富尔塞塔.(*富尔塞塔FursataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fursata",
		TitleName: "富尔塞塔",
		TitleCode: "b_fursata",
	}
}
