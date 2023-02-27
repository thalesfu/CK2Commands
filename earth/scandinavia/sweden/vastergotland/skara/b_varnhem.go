package skara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦恩赫姆VarnhemBarony struct {
	feud.BaseBarony
}

var BVarnhem瓦恩赫姆 feud.Barony = &瓦恩赫姆VarnhemBarony{}

func init() {
    f := BVarnhem瓦恩赫姆.(*瓦恩赫姆VarnhemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "varnhem",
		TitleName: "瓦恩赫姆",
		TitleCode: "b_varnhem",
	}
}
