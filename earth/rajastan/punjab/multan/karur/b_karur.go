package karur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽卢罗KarurBarony struct {
	feud.BaseBarony
}

var BKarur伽卢罗 feud.Barony = &伽卢罗KarurBarony{}

func init() {
	f := BKarur伽卢罗.(*伽卢罗KarurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karur",
		TitleName: "伽卢罗",
		TitleCode: "b_karur",
	}
}
