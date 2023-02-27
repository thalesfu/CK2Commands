package mallabhum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗湿奴城BishnupurBarony struct {
	feud.BaseBarony
}

var BBishnupur毗湿奴城 feud.Barony = &毗湿奴城BishnupurBarony{}

func init() {
    f := BBishnupur毗湿奴城.(*毗湿奴城BishnupurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bishnupur",
		TitleName: "毗湿奴城",
		TitleCode: "b_bishnupur",
	}
}
