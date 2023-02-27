package gloucester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂克斯伯里TewkesburyBarony struct {
	feud.BaseBarony
}

var BTewkesbury蒂克斯伯里 feud.Barony = &蒂克斯伯里TewkesburyBarony{}

func init() {
    f := BTewkesbury蒂克斯伯里.(*蒂克斯伯里TewkesburyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tewkesbury",
		TitleName: "蒂克斯伯里",
		TitleCode: "b_tewkesbury",
	}
}
