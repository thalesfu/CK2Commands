package katehar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 殑伽河门GangadvaraBarony struct {
	feud.BaseBarony
}

var BGangadvara殑伽河门 feud.Barony = &殑伽河门GangadvaraBarony{}

func init() {
	f := BGangadvara殑伽河门.(*殑伽河门GangadvaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gangadvara",
		TitleName: "殑伽河门",
		TitleCode: "b_gangadvara",
	}
}
