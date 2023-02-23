package breisgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗特韦尔RottweilBarony struct {
	feud.BaseBarony
}

var BRottweil罗特韦尔 feud.Barony = &罗特韦尔RottweilBarony{}

func init() {
	f := BRottweil罗特韦尔.(*罗特韦尔RottweilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rottweil",
		TitleName: "罗特韦尔",
		TitleCode: "b_rottweil",
	}
}
