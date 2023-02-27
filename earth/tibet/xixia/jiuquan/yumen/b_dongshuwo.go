package yumen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 东树窝DongshuwoBarony struct {
	feud.BaseBarony
}

var BDongshuwo东树窝 feud.Barony = &东树窝DongshuwoBarony{}

func init() {
    f := BDongshuwo东树窝.(*东树窝DongshuwoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dongshuwo",
		TitleName: "东树窝",
		TitleCode: "b_dongshuwo",
	}
}
