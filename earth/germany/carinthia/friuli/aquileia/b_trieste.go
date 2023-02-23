package aquileia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 的里雅斯特TriesteBarony struct {
	feud.BaseBarony
}

var BTrieste的里雅斯特 feud.Barony = &的里雅斯特TriesteBarony{}

func init() {
	f := BTrieste的里雅斯特.(*的里雅斯特TriesteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trieste",
		TitleName: "的里雅斯特",
		TitleCode: "b_trieste",
	}
}
