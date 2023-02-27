package khilok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 篾儿乞Merkit_khilokBarony struct {
	feud.BaseBarony
}

var BMerkit_khilok篾儿乞 feud.Barony = &篾儿乞Merkit_khilokBarony{}

func init() {
    f := BMerkit_khilok篾儿乞.(*篾儿乞Merkit_khilokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "merkit_khilok",
		TitleName: "篾儿乞",
		TitleCode: "b_merkit_khilok",
	}
}
