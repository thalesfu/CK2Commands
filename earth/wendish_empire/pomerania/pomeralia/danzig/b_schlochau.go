package danzig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施洛豪SchlochauBarony struct {
	feud.BaseBarony
}

var BSchlochau施洛豪 feud.Barony = &施洛豪SchlochauBarony{}

func init() {
    f := BSchlochau施洛豪.(*施洛豪SchlochauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "schlochau",
		TitleName: "施洛豪",
		TitleCode: "b_schlochau",
	}
}
