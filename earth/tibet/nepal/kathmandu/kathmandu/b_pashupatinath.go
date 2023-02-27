package kathmandu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波输波底那他PashupatinathBarony struct {
	feud.BaseBarony
}

var BPashupatinath波输波底那他 feud.Barony = &波输波底那他PashupatinathBarony{}

func init() {
    f := BPashupatinath波输波底那他.(*波输波底那他PashupatinathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pashupatinath",
		TitleName: "波输波底那他",
		TitleCode: "b_pashupatinath",
	}
}
