package tus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法达德次FagdatdezhBarony struct {
	feud.BaseBarony
}

var BFagdatdezh法达德次 feud.Barony = &法达德次FagdatdezhBarony{}

func init() {
    f := BFagdatdezh法达德次.(*法达德次FagdatdezhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fagdatdezh",
		TitleName: "法达德次",
		TitleCode: "b_fagdatdezh",
	}
}
