package ayodhya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿踰陀AyodhyaBarony struct {
	feud.BaseBarony
}

var BAyodhya阿踰陀 feud.Barony = &阿踰陀AyodhyaBarony{}

func init() {
    f := BAyodhya阿踰陀.(*阿踰陀AyodhyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ayodhya",
		TitleName: "阿踰陀",
		TitleCode: "b_ayodhya",
	}
}
