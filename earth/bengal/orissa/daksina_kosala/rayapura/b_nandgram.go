package rayapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 难陀伽罗摩NandgramBarony struct {
	feud.BaseBarony
}

var BNandgram难陀伽罗摩 feud.Barony = &难陀伽罗摩NandgramBarony{}

func init() {
    f := BNandgram难陀伽罗摩.(*难陀伽罗摩NandgramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nandgram",
		TitleName: "难陀伽罗摩",
		TitleCode: "b_nandgram",
	}
}
