package gurjaratra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾里安KharianBarony struct {
	feud.BaseBarony
}

var BKharian贾里安 feud.Barony = &贾里安KharianBarony{}

func init() {
    f := BKharian贾里安.(*贾里安KharianBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kharian",
		TitleName: "贾里安",
		TitleCode: "b_kharian",
	}
}
