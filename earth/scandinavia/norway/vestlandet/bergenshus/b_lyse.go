package bergenshus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕瑟LyseBarony struct {
	feud.BaseBarony
}

var BLyse吕瑟 feud.Barony = &吕瑟LyseBarony{}

func init() {
    f := BLyse吕瑟.(*吕瑟LyseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lyse",
		TitleName: "吕瑟",
		TitleCode: "b_lyse",
	}
}
