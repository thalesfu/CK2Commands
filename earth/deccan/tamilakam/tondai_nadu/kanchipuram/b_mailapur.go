package kanchipuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅罗补罗MailapurBarony struct {
	feud.BaseBarony
}

var BMailapur梅罗补罗 feud.Barony = &梅罗补罗MailapurBarony{}

func init() {
    f := BMailapur梅罗补罗.(*梅罗补罗MailapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mailapur",
		TitleName: "梅罗补罗",
		TitleCode: "b_mailapur",
	}
}
