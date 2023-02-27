package maan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MaanCounty interface {
    feud.County
    BAfra阿夫拉() 	feud.Barony
    BBubeita部贝塔() 	feud.Barony
    BBuseira布赛拉() 	feud.Barony
    BMaan马安() 	feud.Barony
    BMaanalhasa哈赛() 	feud.Barony
    BMaanaljafr杰夫尔() 	feud.Barony
    BMutah穆阿泰() 	feud.Barony
    BShubak苏巴克() 	feud.Barony
}

type 马安MaanCounty struct {
	feud.BaseCounty
	阿夫拉Afra 	feud.Barony
	部贝塔Bubeita 	feud.Barony
	布赛拉Buseira 	feud.Barony
	马安Maan 	feud.Barony
	哈赛Maanalhasa 	feud.Barony
	杰夫尔Maanaljafr 	feud.Barony
	穆阿泰Mutah 	feud.Barony
	苏巴克Shubak 	feud.Barony
}

func (c *马安MaanCounty) BAfra阿夫拉() feud.Barony {
	return c.阿夫拉Afra
}
    
func (c *马安MaanCounty) BBubeita部贝塔() feud.Barony {
	return c.部贝塔Bubeita
}
    
func (c *马安MaanCounty) BBuseira布赛拉() feud.Barony {
	return c.布赛拉Buseira
}
    
func (c *马安MaanCounty) BMaan马安() feud.Barony {
	return c.马安Maan
}
    
func (c *马安MaanCounty) BMaanalhasa哈赛() feud.Barony {
	return c.哈赛Maanalhasa
}
    
func (c *马安MaanCounty) BMaanaljafr杰夫尔() feud.Barony {
	return c.杰夫尔Maanaljafr
}
    
func (c *马安MaanCounty) BMutah穆阿泰() feud.Barony {
	return c.穆阿泰Mutah
}
    
func (c *马安MaanCounty) BShubak苏巴克() feud.Barony {
	return c.苏巴克Shubak
}
    
var CMaan马安 MaanCounty = &马安MaanCounty{}

func init() {
	f := CMaan马安.(*马安MaanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "783",
		Title:     "maan",
		TitleName: "马安",
		TitleCode: "c_maan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿夫拉Afra = BAfra阿夫拉
	f.阿夫拉Afra.SetParent(f)
	
	f.部贝塔Bubeita = BBubeita部贝塔
	f.部贝塔Bubeita.SetParent(f)
	
	f.布赛拉Buseira = BBuseira布赛拉
	f.布赛拉Buseira.SetParent(f)
	
	f.马安Maan = BMaan马安
	f.马安Maan.SetParent(f)
	
	f.哈赛Maanalhasa = BMaanalhasa哈赛
	f.哈赛Maanalhasa.SetParent(f)
	
	f.杰夫尔Maanaljafr = BMaanaljafr杰夫尔
	f.杰夫尔Maanaljafr.SetParent(f)
	
	f.穆阿泰Mutah = BMutah穆阿泰
	f.穆阿泰Mutah.SetParent(f)
	
	f.苏巴克Shubak = BShubak苏巴克
	f.苏巴克Shubak.SetParent(f)
	
}
