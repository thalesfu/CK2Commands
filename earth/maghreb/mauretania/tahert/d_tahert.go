package tahert

import (
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/tahert/ashir"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/tahert/atlas_mnt"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/tahert/hanyan"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/tahert/lemdiyya"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TahertDuke interface {
    feud.Duke
    CAshir阿希尔() 	ashir.AshirCounty
    CAtlas_mnt阿斯克达尔() 	atlas_mnt.Atlas_mntCounty
    CHanyan亚拉拉() 	hanyan.HanyanCounty
    CLemdiyya提亚雷特() 	lemdiyya.LemdiyyaCounty
}

type 提亚雷特TahertDuke struct {
	feud.BaseDuke
	阿希尔Ashir 	ashir.AshirCounty
	阿斯克达尔Atlas_mnt 	atlas_mnt.Atlas_mntCounty
	亚拉拉Hanyan 	hanyan.HanyanCounty
	提亚雷特Lemdiyya 	lemdiyya.LemdiyyaCounty
}

func (d *提亚雷特TahertDuke) CAshir阿希尔() ashir.AshirCounty {
	return d.阿希尔Ashir
}
    
func (d *提亚雷特TahertDuke) CAtlas_mnt阿斯克达尔() atlas_mnt.Atlas_mntCounty {
	return d.阿斯克达尔Atlas_mnt
}
    
func (d *提亚雷特TahertDuke) CHanyan亚拉拉() hanyan.HanyanCounty {
	return d.亚拉拉Hanyan
}
    
func (d *提亚雷特TahertDuke) CLemdiyya提亚雷特() lemdiyya.LemdiyyaCounty {
	return d.提亚雷特Lemdiyya
}
    
var DTahert提亚雷特 TahertDuke = &提亚雷特TahertDuke{}

func init() {
	f := DTahert提亚雷特.(*提亚雷特TahertDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "tahert",
		TitleName: "提亚雷特",
		TitleCode: "d_tahert",
		Counties:  map[string]feud.County{},
	}

	f.阿希尔Ashir = ashir.CAshir阿希尔
	f.阿希尔Ashir.SetParent(f)
	
	f.阿斯克达尔Atlas_mnt = atlas_mnt.CAtlas_mnt阿斯克达尔
	f.阿斯克达尔Atlas_mnt.SetParent(f)
	
	f.亚拉拉Hanyan = hanyan.CHanyan亚拉拉
	f.亚拉拉Hanyan.SetParent(f)
	
	f.提亚雷特Lemdiyya = lemdiyya.CLemdiyya提亚雷特
	f.提亚雷特Lemdiyya.SetParent(f)
	
}
