package kuznetsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乔尔诺耶Chernoye_kuznetskBarony struct {
	feud.BaseBarony
}

var BChernoye_kuznetsk乔尔诺耶 feud.Barony = &乔尔诺耶Chernoye_kuznetskBarony{}

func init() {
    f := BChernoye_kuznetsk乔尔诺耶.(*乔尔诺耶Chernoye_kuznetskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chernoye_kuznetsk",
		TitleName: "乔尔诺耶",
		TitleCode: "b_chernoye_kuznetsk",
	}
}
