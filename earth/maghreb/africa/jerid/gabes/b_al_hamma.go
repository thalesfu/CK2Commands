package gabes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈迈Al_hammaBarony struct {
	feud.BaseBarony
}

var BAl_hamma哈迈 feud.Barony = &哈迈Al_hammaBarony{}

func init() {
    f := BAl_hamma哈迈.(*哈迈Al_hammaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_hamma",
		TitleName: "哈迈",
		TitleCode: "b_al_hamma",
	}
}
