package gurvan_saikhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Gurvan_saikhanCounty interface {
    feud.County
    BBaruun_saikhany_nuruu巴伦() 	feud.Barony
    BBayandalai巴彦达赖() 	feud.Barony
    BDund_saikhany_nuruu敦德() 	feud.Barony
    BGurvan_saikhan古尔班赛罕() 	feud.Barony
    BKhankhongor汗洪戈尔() 	feud.Barony
    BKhurmen呼尔门() 	feud.Barony
    BZuun_saikhany_nuruu准() 	feud.Barony
}

type 古尔班赛罕Gurvan_saikhanCounty struct {
	feud.BaseCounty
	巴伦Baruun_saikhany_nuruu 	feud.Barony
	巴彦达赖Bayandalai 	feud.Barony
	敦德Dund_saikhany_nuruu 	feud.Barony
	古尔班赛罕Gurvan_saikhan 	feud.Barony
	汗洪戈尔Khankhongor 	feud.Barony
	呼尔门Khurmen 	feud.Barony
	准Zuun_saikhany_nuruu 	feud.Barony
}

func (c *古尔班赛罕Gurvan_saikhanCounty) BBaruun_saikhany_nuruu巴伦() feud.Barony {
	return c.巴伦Baruun_saikhany_nuruu
}
    
func (c *古尔班赛罕Gurvan_saikhanCounty) BBayandalai巴彦达赖() feud.Barony {
	return c.巴彦达赖Bayandalai
}
    
func (c *古尔班赛罕Gurvan_saikhanCounty) BDund_saikhany_nuruu敦德() feud.Barony {
	return c.敦德Dund_saikhany_nuruu
}
    
func (c *古尔班赛罕Gurvan_saikhanCounty) BGurvan_saikhan古尔班赛罕() feud.Barony {
	return c.古尔班赛罕Gurvan_saikhan
}
    
func (c *古尔班赛罕Gurvan_saikhanCounty) BKhankhongor汗洪戈尔() feud.Barony {
	return c.汗洪戈尔Khankhongor
}
    
func (c *古尔班赛罕Gurvan_saikhanCounty) BKhurmen呼尔门() feud.Barony {
	return c.呼尔门Khurmen
}
    
func (c *古尔班赛罕Gurvan_saikhanCounty) BZuun_saikhany_nuruu准() feud.Barony {
	return c.准Zuun_saikhany_nuruu
}
    
var CGurvan_saikhan古尔班赛罕 Gurvan_saikhanCounty = &古尔班赛罕Gurvan_saikhanCounty{}

func init() {
	f := CGurvan_saikhan古尔班赛罕.(*古尔班赛罕Gurvan_saikhanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1917",
		Title:     "gurvan_saikhan",
		TitleName: "古尔班赛罕",
		TitleCode: "c_gurvan_saikhan",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴伦Baruun_saikhany_nuruu = BBaruun_saikhany_nuruu巴伦
	f.巴伦Baruun_saikhany_nuruu.SetParent(f)
	
	f.巴彦达赖Bayandalai = BBayandalai巴彦达赖
	f.巴彦达赖Bayandalai.SetParent(f)
	
	f.敦德Dund_saikhany_nuruu = BDund_saikhany_nuruu敦德
	f.敦德Dund_saikhany_nuruu.SetParent(f)
	
	f.古尔班赛罕Gurvan_saikhan = BGurvan_saikhan古尔班赛罕
	f.古尔班赛罕Gurvan_saikhan.SetParent(f)
	
	f.汗洪戈尔Khankhongor = BKhankhongor汗洪戈尔
	f.汗洪戈尔Khankhongor.SetParent(f)
	
	f.呼尔门Khurmen = BKhurmen呼尔门
	f.呼尔门Khurmen.SetParent(f)
	
	f.准Zuun_saikhany_nuruu = BZuun_saikhany_nuruu准
	f.准Zuun_saikhany_nuruu.SetParent(f)
	
}
