package asni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AsniCounty interface {
    feud.County
    BBhitargaon内村() 	feud.Barony
    BBindki频陀吉() 	feud.Barony
    BKhaga卡迦() 	feud.Barony
    BKora厥罗() 	feud.Barony
    BMusanagar牟沙那揭罗() 	feud.Barony
    BShivrajpur湿婆罗阇补罗() 	feud.Barony
}

type 阿斯尼AsniCounty struct {
	feud.BaseCounty
	内村Bhitargaon 	feud.Barony
	频陀吉Bindki 	feud.Barony
	卡迦Khaga 	feud.Barony
	厥罗Kora 	feud.Barony
	牟沙那揭罗Musanagar 	feud.Barony
	湿婆罗阇补罗Shivrajpur 	feud.Barony
}

func (c *阿斯尼AsniCounty) BBhitargaon内村() feud.Barony {
	return c.内村Bhitargaon
}
    
func (c *阿斯尼AsniCounty) BBindki频陀吉() feud.Barony {
	return c.频陀吉Bindki
}
    
func (c *阿斯尼AsniCounty) BKhaga卡迦() feud.Barony {
	return c.卡迦Khaga
}
    
func (c *阿斯尼AsniCounty) BKora厥罗() feud.Barony {
	return c.厥罗Kora
}
    
func (c *阿斯尼AsniCounty) BMusanagar牟沙那揭罗() feud.Barony {
	return c.牟沙那揭罗Musanagar
}
    
func (c *阿斯尼AsniCounty) BShivrajpur湿婆罗阇补罗() feud.Barony {
	return c.湿婆罗阇补罗Shivrajpur
}
    
var CAsni阿斯尼 AsniCounty = &阿斯尼AsniCounty{}

func init() {
	f := CAsni阿斯尼.(*阿斯尼AsniCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1283",
		Title:     "asni",
		TitleName: "阿斯尼",
		TitleCode: "c_asni",
		Baronies:  map[string]feud.Barony{},
	}

	f.内村Bhitargaon = BBhitargaon内村
	f.内村Bhitargaon.SetParent(f)
	
	f.频陀吉Bindki = BBindki频陀吉
	f.频陀吉Bindki.SetParent(f)
	
	f.卡迦Khaga = BKhaga卡迦
	f.卡迦Khaga.SetParent(f)
	
	f.厥罗Kora = BKora厥罗
	f.厥罗Kora.SetParent(f)
	
	f.牟沙那揭罗Musanagar = BMusanagar牟沙那揭罗
	f.牟沙那揭罗Musanagar.SetParent(f)
	
	f.湿婆罗阇补罗Shivrajpur = BShivrajpur湿婆罗阇补罗
	f.湿婆罗阇补罗Shivrajpur.SetParent(f)
	
}
