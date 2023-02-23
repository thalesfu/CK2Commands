package gondar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡德尔GondarBarony struct {
	feud.BaseBarony
}

var BGondar贡德尔 feud.Barony = &贡德尔GondarBarony{}

func init() {
	f := BGondar贡德尔.(*贡德尔GondarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gondar",
		TitleName: "贡德尔",
		TitleCode: "b_gondar",
	}
}
