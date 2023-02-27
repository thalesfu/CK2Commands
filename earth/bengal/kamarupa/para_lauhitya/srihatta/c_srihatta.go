package srihatta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SrihattaCounty interface {
    feud.County
    BAzampur阿赡补罗() 	feud.Barony
    BEgarosindur夷迦卢斯那突() 	feud.Barony
    BHabiganj诃毗犍阇() 	feud.Barony
    BJangalbari旬迦罗婆利() 	feud.Barony
    BKunnau贡奴() 	feud.Barony
    BNasirabad那赐罗跋() 	feud.Barony
    BSrihatta室利诃吒() 	feud.Barony
}

type 室利诃吒SrihattaCounty struct {
	feud.BaseCounty
	阿赡补罗Azampur 	feud.Barony
	夷迦卢斯那突Egarosindur 	feud.Barony
	诃毗犍阇Habiganj 	feud.Barony
	旬迦罗婆利Jangalbari 	feud.Barony
	贡奴Kunnau 	feud.Barony
	那赐罗跋Nasirabad 	feud.Barony
	室利诃吒Srihatta 	feud.Barony
}

func (c *室利诃吒SrihattaCounty) BAzampur阿赡补罗() feud.Barony {
	return c.阿赡补罗Azampur
}
    
func (c *室利诃吒SrihattaCounty) BEgarosindur夷迦卢斯那突() feud.Barony {
	return c.夷迦卢斯那突Egarosindur
}
    
func (c *室利诃吒SrihattaCounty) BHabiganj诃毗犍阇() feud.Barony {
	return c.诃毗犍阇Habiganj
}
    
func (c *室利诃吒SrihattaCounty) BJangalbari旬迦罗婆利() feud.Barony {
	return c.旬迦罗婆利Jangalbari
}
    
func (c *室利诃吒SrihattaCounty) BKunnau贡奴() feud.Barony {
	return c.贡奴Kunnau
}
    
func (c *室利诃吒SrihattaCounty) BNasirabad那赐罗跋() feud.Barony {
	return c.那赐罗跋Nasirabad
}
    
func (c *室利诃吒SrihattaCounty) BSrihatta室利诃吒() feud.Barony {
	return c.室利诃吒Srihatta
}
    
var CSrihatta室利诃吒 SrihattaCounty = &室利诃吒SrihattaCounty{}

func init() {
	f := CSrihatta室利诃吒.(*室利诃吒SrihattaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1245",
		Title:     "srihatta",
		TitleName: "室利诃吒",
		TitleCode: "c_srihatta",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿赡补罗Azampur = BAzampur阿赡补罗
	f.阿赡补罗Azampur.SetParent(f)
	
	f.夷迦卢斯那突Egarosindur = BEgarosindur夷迦卢斯那突
	f.夷迦卢斯那突Egarosindur.SetParent(f)
	
	f.诃毗犍阇Habiganj = BHabiganj诃毗犍阇
	f.诃毗犍阇Habiganj.SetParent(f)
	
	f.旬迦罗婆利Jangalbari = BJangalbari旬迦罗婆利
	f.旬迦罗婆利Jangalbari.SetParent(f)
	
	f.贡奴Kunnau = BKunnau贡奴
	f.贡奴Kunnau.SetParent(f)
	
	f.那赐罗跋Nasirabad = BNasirabad那赐罗跋
	f.那赐罗跋Nasirabad.SetParent(f)
	
	f.室利诃吒Srihatta = BSrihatta室利诃吒
	f.室利诃吒Srihatta.SetParent(f)
	
}
