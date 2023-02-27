package canda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳吉诃德NagbhidBarony struct {
	feud.BaseBarony
}

var BNagbhid纳吉诃德 feud.Barony = &纳吉诃德NagbhidBarony{}

func init() {
    f := BNagbhid纳吉诃德.(*纳吉诃德NagbhidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagbhid",
		TitleName: "纳吉诃德",
		TitleCode: "b_nagbhid",
	}
}
