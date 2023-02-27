package kantor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KantorCounty interface {
    feud.County
    BBadogo巴多戈() 	feud.Barony
    BBidangala比当加拉() 	feud.Barony
    BGoudiourou古迪乌鲁() 	feud.Barony
    BKamindalah卡明达拉() 	feud.Barony
    BKantor坎托尔() 	feud.Barony
    BMandio芒迪奥() 	feud.Barony
    BOuli乌利() 	feud.Barony
    BSoliko索利科() 	feud.Barony
}

type 坎托尔KantorCounty struct {
	feud.BaseCounty
	巴多戈Badogo 	feud.Barony
	比当加拉Bidangala 	feud.Barony
	古迪乌鲁Goudiourou 	feud.Barony
	卡明达拉Kamindalah 	feud.Barony
	坎托尔Kantor 	feud.Barony
	芒迪奥Mandio 	feud.Barony
	乌利Ouli 	feud.Barony
	索利科Soliko 	feud.Barony
}

func (c *坎托尔KantorCounty) BBadogo巴多戈() feud.Barony {
	return c.巴多戈Badogo
}
    
func (c *坎托尔KantorCounty) BBidangala比当加拉() feud.Barony {
	return c.比当加拉Bidangala
}
    
func (c *坎托尔KantorCounty) BGoudiourou古迪乌鲁() feud.Barony {
	return c.古迪乌鲁Goudiourou
}
    
func (c *坎托尔KantorCounty) BKamindalah卡明达拉() feud.Barony {
	return c.卡明达拉Kamindalah
}
    
func (c *坎托尔KantorCounty) BKantor坎托尔() feud.Barony {
	return c.坎托尔Kantor
}
    
func (c *坎托尔KantorCounty) BMandio芒迪奥() feud.Barony {
	return c.芒迪奥Mandio
}
    
func (c *坎托尔KantorCounty) BOuli乌利() feud.Barony {
	return c.乌利Ouli
}
    
func (c *坎托尔KantorCounty) BSoliko索利科() feud.Barony {
	return c.索利科Soliko
}
    
var CKantor坎托尔 KantorCounty = &坎托尔KantorCounty{}

func init() {
	f := CKantor坎托尔.(*坎托尔KantorCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1765",
		Title:     "kantor",
		TitleName: "坎托尔",
		TitleCode: "c_kantor",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴多戈Badogo = BBadogo巴多戈
	f.巴多戈Badogo.SetParent(f)
	
	f.比当加拉Bidangala = BBidangala比当加拉
	f.比当加拉Bidangala.SetParent(f)
	
	f.古迪乌鲁Goudiourou = BGoudiourou古迪乌鲁
	f.古迪乌鲁Goudiourou.SetParent(f)
	
	f.卡明达拉Kamindalah = BKamindalah卡明达拉
	f.卡明达拉Kamindalah.SetParent(f)
	
	f.坎托尔Kantor = BKantor坎托尔
	f.坎托尔Kantor.SetParent(f)
	
	f.芒迪奥Mandio = BMandio芒迪奥
	f.芒迪奥Mandio.SetParent(f)
	
	f.乌利Ouli = BOuli乌利
	f.乌利Ouli.SetParent(f)
	
	f.索利科Soliko = BSoliko索利科
	f.索利科Soliko.SetParent(f)
	
}
