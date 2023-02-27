package bezhetsky_verh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柳班LyubanBarony struct {
	feud.BaseBarony
}

var BLyuban柳班 feud.Barony = &柳班LyubanBarony{}

func init() {
    f := BLyuban柳班.(*柳班LyubanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lyuban",
		TitleName: "柳班",
		TitleCode: "b_lyuban",
	}
}
