package kazan

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/kazan/alabuga"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/kazan/khlynov"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/kazan/kremenchuk"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/kazan/qazan"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/kazan/vetluga"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KazanDuke interface {
    feud.Duke
    CAlabuga阿拉布加() 	alabuga.AlabugaCounty
    CKhlynov赫雷诺夫() 	khlynov.KhlynovCounty
    CKremenchuk克列缅丘格() 	kremenchuk.KremenchukCounty
    CQazan喀山() 	qazan.QazanCounty
    CVetluga韦特卢加() 	vetluga.VetlugaCounty
}

type 喀山KazanDuke struct {
	feud.BaseDuke
	阿拉布加Alabuga 	alabuga.AlabugaCounty
	赫雷诺夫Khlynov 	khlynov.KhlynovCounty
	克列缅丘格Kremenchuk 	kremenchuk.KremenchukCounty
	喀山Qazan 	qazan.QazanCounty
	韦特卢加Vetluga 	vetluga.VetlugaCounty
}

func (d *喀山KazanDuke) CAlabuga阿拉布加() alabuga.AlabugaCounty {
	return d.阿拉布加Alabuga
}
    
func (d *喀山KazanDuke) CKhlynov赫雷诺夫() khlynov.KhlynovCounty {
	return d.赫雷诺夫Khlynov
}
    
func (d *喀山KazanDuke) CKremenchuk克列缅丘格() kremenchuk.KremenchukCounty {
	return d.克列缅丘格Kremenchuk
}
    
func (d *喀山KazanDuke) CQazan喀山() qazan.QazanCounty {
	return d.喀山Qazan
}
    
func (d *喀山KazanDuke) CVetluga韦特卢加() vetluga.VetlugaCounty {
	return d.韦特卢加Vetluga
}
    
var DKazan喀山 KazanDuke = &喀山KazanDuke{}

func init() {
	f := DKazan喀山.(*喀山KazanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kazan",
		TitleName: "喀山",
		TitleCode: "d_kazan",
		Counties:  map[string]feud.County{},
	}

	f.阿拉布加Alabuga = alabuga.CAlabuga阿拉布加
	f.阿拉布加Alabuga.SetParent(f)
	
	f.赫雷诺夫Khlynov = khlynov.CKhlynov赫雷诺夫
	f.赫雷诺夫Khlynov.SetParent(f)
	
	f.克列缅丘格Kremenchuk = kremenchuk.CKremenchuk克列缅丘格
	f.克列缅丘格Kremenchuk.SetParent(f)
	
	f.喀山Qazan = qazan.CQazan喀山
	f.喀山Qazan.SetParent(f)
	
	f.韦特卢加Vetluga = vetluga.CVetluga韦特卢加
	f.韦特卢加Vetluga.SetParent(f)
	
}
