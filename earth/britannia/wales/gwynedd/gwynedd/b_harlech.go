package gwynedd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈勒赫HarlechBarony struct {
	feud.BaseBarony
}

var BHarlech哈勒赫 feud.Barony = &哈勒赫HarlechBarony{}

func init() {
	f := BHarlech哈勒赫.(*哈勒赫HarlechBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harlech",
		TitleName: "哈勒赫",
		TitleCode: "b_harlech",
	}
}
