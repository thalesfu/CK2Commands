package amaravati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗奴军荼VinukondaBarony struct {
	feud.BaseBarony
}

var BVinukonda毗奴军荼 feud.Barony = &毗奴军荼VinukondaBarony{}

func init() {
    f := BVinukonda毗奴军荼.(*毗奴军荼VinukondaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vinukonda",
		TitleName: "毗奴军荼",
		TitleCode: "b_vinukonda",
	}
}
