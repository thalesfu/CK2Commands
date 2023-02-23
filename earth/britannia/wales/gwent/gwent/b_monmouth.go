package gwent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙茅斯MonmouthBarony struct {
	feud.BaseBarony
}

var BMonmouth蒙茅斯 feud.Barony = &蒙茅斯MonmouthBarony{}

func init() {
	f := BMonmouth蒙茅斯.(*蒙茅斯MonmouthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monmouth",
		TitleName: "蒙茅斯",
		TitleCode: "b_monmouth",
	}
}
