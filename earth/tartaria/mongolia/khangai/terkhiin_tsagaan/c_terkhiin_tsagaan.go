package terkhiin_tsagaan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Terkhiin_tsagaanCounty interface {
    feud.County
    BAvgaldai阿布噶勒岱() 	feud.Barony
    BKhadat哈达特() 	feud.Barony
    BKhorgo霍尔戈() 	feud.Barony
    BKhunt洪特() 	feud.Barony
    BTariat塔里亚特() 	feud.Barony
    BTerkhiin_tsagaan特尔欣查干() 	feud.Barony
    BTsahir查黑尔() 	feud.Barony
}

type 特尔欣查干Terkhiin_tsagaanCounty struct {
	feud.BaseCounty
	阿布噶勒岱Avgaldai 	feud.Barony
	哈达特Khadat 	feud.Barony
	霍尔戈Khorgo 	feud.Barony
	洪特Khunt 	feud.Barony
	塔里亚特Tariat 	feud.Barony
	特尔欣查干Terkhiin_tsagaan 	feud.Barony
	查黑尔Tsahir 	feud.Barony
}

func (c *特尔欣查干Terkhiin_tsagaanCounty) BAvgaldai阿布噶勒岱() feud.Barony {
	return c.阿布噶勒岱Avgaldai
}
    
func (c *特尔欣查干Terkhiin_tsagaanCounty) BKhadat哈达特() feud.Barony {
	return c.哈达特Khadat
}
    
func (c *特尔欣查干Terkhiin_tsagaanCounty) BKhorgo霍尔戈() feud.Barony {
	return c.霍尔戈Khorgo
}
    
func (c *特尔欣查干Terkhiin_tsagaanCounty) BKhunt洪特() feud.Barony {
	return c.洪特Khunt
}
    
func (c *特尔欣查干Terkhiin_tsagaanCounty) BTariat塔里亚特() feud.Barony {
	return c.塔里亚特Tariat
}
    
func (c *特尔欣查干Terkhiin_tsagaanCounty) BTerkhiin_tsagaan特尔欣查干() feud.Barony {
	return c.特尔欣查干Terkhiin_tsagaan
}
    
func (c *特尔欣查干Terkhiin_tsagaanCounty) BTsahir查黑尔() feud.Barony {
	return c.查黑尔Tsahir
}
    
var CTerkhiin_tsagaan特尔欣查干 Terkhiin_tsagaanCounty = &特尔欣查干Terkhiin_tsagaanCounty{}

func init() {
	f := CTerkhiin_tsagaan特尔欣查干.(*特尔欣查干Terkhiin_tsagaanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1915",
		Title:     "terkhiin_tsagaan",
		TitleName: "特尔欣查干",
		TitleCode: "c_terkhiin_tsagaan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿布噶勒岱Avgaldai = BAvgaldai阿布噶勒岱
	f.阿布噶勒岱Avgaldai.SetParent(f)
	
	f.哈达特Khadat = BKhadat哈达特
	f.哈达特Khadat.SetParent(f)
	
	f.霍尔戈Khorgo = BKhorgo霍尔戈
	f.霍尔戈Khorgo.SetParent(f)
	
	f.洪特Khunt = BKhunt洪特
	f.洪特Khunt.SetParent(f)
	
	f.塔里亚特Tariat = BTariat塔里亚特
	f.塔里亚特Tariat.SetParent(f)
	
	f.特尔欣查干Terkhiin_tsagaan = BTerkhiin_tsagaan特尔欣查干
	f.特尔欣查干Terkhiin_tsagaan.SetParent(f)
	
	f.查黑尔Tsahir = BTsahir查黑尔
	f.查黑尔Tsahir.SetParent(f)
	
}
