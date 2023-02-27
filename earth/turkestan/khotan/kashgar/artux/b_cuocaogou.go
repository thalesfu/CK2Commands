package artux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 错草沟CuocaogouBarony struct {
	feud.BaseBarony
}

var BCuocaogou错草沟 feud.Barony = &错草沟CuocaogouBarony{}

func init() {
    f := BCuocaogou错草沟.(*错草沟CuocaogouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cuocaogou",
		TitleName: "错草沟",
		TitleCode: "b_cuocaogou",
	}
}
