package pereyaslavl_zalessky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库潘斯科耶KupanskoyeBarony struct {
	feud.BaseBarony
}

var BKupanskoye库潘斯科耶 feud.Barony = &库潘斯科耶KupanskoyeBarony{}

func init() {
    f := BKupanskoye库潘斯科耶.(*库潘斯科耶KupanskoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kupanskoye",
		TitleName: "库潘斯科耶",
		TitleCode: "b_kupanskoye",
	}
}
