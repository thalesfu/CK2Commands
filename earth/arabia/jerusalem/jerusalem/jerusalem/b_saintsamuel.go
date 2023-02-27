package jerusalem

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣撒慕尔SaintsamuelBarony struct {
	feud.BaseBarony
}

var BSaintsamuel圣撒慕尔 feud.Barony = &圣撒慕尔SaintsamuelBarony{}

func init() {
    f := BSaintsamuel圣撒慕尔.(*圣撒慕尔SaintsamuelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saintsamuel",
		TitleName: "圣撒慕尔",
		TitleCode: "b_saintsamuel",
	}
}
