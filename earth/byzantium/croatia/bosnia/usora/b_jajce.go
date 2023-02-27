package usora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚伊采JajceBarony struct {
	feud.BaseBarony
}

var BJajce亚伊采 feud.Barony = &亚伊采JajceBarony{}

func init() {
    f := BJajce亚伊采.(*亚伊采JajceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jajce",
		TitleName: "亚伊采",
		TitleCode: "b_jajce",
	}
}
