package bashkirs

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/bashkirs/bashkirs"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/bashkirs/belaya"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BashkirsDuke interface {
    feud.Duke
    CBashkirs巴什基里亚() 	bashkirs.BashkirsCounty
    CBelaya别拉亚() 	belaya.BelayaCounty
}

type 巴什基里亚BashkirsDuke struct {
	feud.BaseDuke
	巴什基里亚Bashkirs 	bashkirs.BashkirsCounty
	别拉亚Belaya 	belaya.BelayaCounty
}

func (d *巴什基里亚BashkirsDuke) CBashkirs巴什基里亚() bashkirs.BashkirsCounty {
	return d.巴什基里亚Bashkirs
}
    
func (d *巴什基里亚BashkirsDuke) CBelaya别拉亚() belaya.BelayaCounty {
	return d.别拉亚Belaya
}
    
var DBashkirs巴什基里亚 BashkirsDuke = &巴什基里亚BashkirsDuke{}

func init() {
	f := DBashkirs巴什基里亚.(*巴什基里亚BashkirsDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "bashkirs",
		TitleName: "巴什基里亚",
		TitleCode: "d_bashkirs",
		Counties:  map[string]feud.County{},
	}

	f.巴什基里亚Bashkirs = bashkirs.CBashkirs巴什基里亚
	f.巴什基里亚Bashkirs.SetParent(f)
	
	f.别拉亚Belaya = belaya.CBelaya别拉亚
	f.别拉亚Belaya.SetParent(f)
	
}
