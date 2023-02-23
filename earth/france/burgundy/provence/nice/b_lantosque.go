package nice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朗托斯克LantosqueBarony struct {
	feud.BaseBarony
}

var BLantosque朗托斯克 feud.Barony = &朗托斯克LantosqueBarony{}

func init() {
	f := BLantosque朗托斯克.(*朗托斯克LantosqueBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lantosque",
		TitleName: "朗托斯克",
		TitleCode: "b_lantosque",
	}
}
