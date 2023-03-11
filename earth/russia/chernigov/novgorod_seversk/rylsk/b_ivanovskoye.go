package rylsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊万诺夫斯科耶IvanovskoyeBarony struct {
	feud.BaseBarony
}

var BIvanovskoye伊万诺夫斯科耶 feud.Barony = &伊万诺夫斯科耶IvanovskoyeBarony{}

func init() {
    f := BIvanovskoye伊万诺夫斯科耶.(*伊万诺夫斯科耶IvanovskoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ivanovskoye",
		TitleName: "伊万诺夫斯科耶",
		TitleCode: "b_ivanovskoye",
	}
}
