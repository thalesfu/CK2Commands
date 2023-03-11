package dihistan

import (
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/dihistan/dihistan"
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/dihistan/kara_kum"
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/dihistan/kyzyl_su"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DihistanDuke interface {
    feud.Duke
    CDihistan大益斯坦() 	dihistan.DihistanCounty
    CKara_kum法拉瓦() 	kara_kum.Kara_kumCounty
    CKyzyl_su克孜勒苏() 	kyzyl_su.Kyzyl_suCounty
}

type 大益斯坦DihistanDuke struct {
	feud.BaseDuke
	大益斯坦Dihistan 	dihistan.DihistanCounty
	法拉瓦Kara_kum 	kara_kum.Kara_kumCounty
	克孜勒苏Kyzyl_su 	kyzyl_su.Kyzyl_suCounty
}

func (d *大益斯坦DihistanDuke) CDihistan大益斯坦() dihistan.DihistanCounty {
	return d.大益斯坦Dihistan
}
    
func (d *大益斯坦DihistanDuke) CKara_kum法拉瓦() kara_kum.Kara_kumCounty {
	return d.法拉瓦Kara_kum
}
    
func (d *大益斯坦DihistanDuke) CKyzyl_su克孜勒苏() kyzyl_su.Kyzyl_suCounty {
	return d.克孜勒苏Kyzyl_su
}
    
var DDihistan大益斯坦 DihistanDuke = &大益斯坦DihistanDuke{}

func init() {
	f := DDihistan大益斯坦.(*大益斯坦DihistanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "dihistan",
		TitleName: "大益斯坦",
		TitleCode: "d_dihistan",
		Counties:  map[string]feud.County{},
	}

	f.大益斯坦Dihistan = dihistan.CDihistan大益斯坦
	f.大益斯坦Dihistan.SetParent(f)
	
	f.法拉瓦Kara_kum = kara_kum.CKara_kum法拉瓦
	f.法拉瓦Kara_kum.SetParent(f)
	
	f.克孜勒苏Kyzyl_su = kyzyl_su.CKyzyl_su克孜勒苏
	f.克孜勒苏Kyzyl_su.SetParent(f)
	
}
