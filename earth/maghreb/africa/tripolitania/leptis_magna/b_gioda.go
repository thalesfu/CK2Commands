package leptis_magna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉欧德GiodaBarony struct {
	feud.BaseBarony
}

var BGioda吉欧德 feud.Barony = &吉欧德GiodaBarony{}

func init() {
    f := BGioda吉欧德.(*吉欧德GiodaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gioda",
		TitleName: "吉欧德",
		TitleCode: "b_gioda",
	}
}
