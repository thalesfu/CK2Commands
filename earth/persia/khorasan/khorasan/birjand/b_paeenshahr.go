package birjand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕恩_沙赫尔PaeenshahrBarony struct {
	feud.BaseBarony
}

var BPaeenshahr帕恩_沙赫尔 feud.Barony = &帕恩_沙赫尔PaeenshahrBarony{}

func init() {
	f := BPaeenshahr帕恩_沙赫尔.(*帕恩_沙赫尔PaeenshahrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paeenshahr",
		TitleName: "帕恩_沙赫尔",
		TitleCode: "b_paeenshahr",
	}
}
