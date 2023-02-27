package pronsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利哈列夫斯科耶戈罗季希Likharevskoye_gorodishchBarony struct {
	feud.BaseBarony
}

var BLikharevskoye_gorodishch利哈列夫斯科耶戈罗季希 feud.Barony = &利哈列夫斯科耶戈罗季希Likharevskoye_gorodishchBarony{}

func init() {
    f := BLikharevskoye_gorodishch利哈列夫斯科耶戈罗季希.(*利哈列夫斯科耶戈罗季希Likharevskoye_gorodishchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "likharevskoye_gorodishch",
		TitleName: "利哈列夫斯科耶戈罗季希",
		TitleCode: "b_likharevskoye_gorodishch",
	}
}
