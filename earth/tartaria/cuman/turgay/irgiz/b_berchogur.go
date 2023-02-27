package irgiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别尔乔古尔BerchogurBarony struct {
	feud.BaseBarony
}

var BBerchogur别尔乔古尔 feud.Barony = &别尔乔古尔BerchogurBarony{}

func init() {
    f := BBerchogur别尔乔古尔.(*别尔乔古尔BerchogurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berchogur",
		TitleName: "别尔乔古尔",
		TitleCode: "b_berchogur",
	}
}
