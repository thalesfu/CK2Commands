package tuggurt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂梅利林TimelilineBarony struct {
	feud.BaseBarony
}

var BTimeliline蒂梅利林 feud.Barony = &蒂梅利林TimelilineBarony{}

func init() {
	f := BTimeliline蒂梅利林.(*蒂梅利林TimelilineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "timeliline",
		TitleName: "蒂梅利林",
		TitleCode: "b_timeliline",
	}
}
