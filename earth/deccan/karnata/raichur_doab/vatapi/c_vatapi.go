package vatapi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VatapiCounty interface {
    feud.County
    BAihole阿赫勒() 	feud.Barony
    BBagalkot巴加尔科特() 	feud.Barony
    BDharwar塔尔瓦尔() 	feud.Barony
    BHalasi诃罗斯() 	feud.Barony
    BParasgad帕拉斯迦得() 	feud.Barony
    BVatapi伐达比() 	feud.Barony
    BVenugrama鞞纽伽罗摩() 	feud.Barony
}

type 伐达比VatapiCounty struct {
	feud.BaseCounty
	阿赫勒Aihole 	feud.Barony
	巴加尔科特Bagalkot 	feud.Barony
	塔尔瓦尔Dharwar 	feud.Barony
	诃罗斯Halasi 	feud.Barony
	帕拉斯迦得Parasgad 	feud.Barony
	伐达比Vatapi 	feud.Barony
	鞞纽伽罗摩Venugrama 	feud.Barony
}

func (c *伐达比VatapiCounty) BAihole阿赫勒() feud.Barony {
	return c.阿赫勒Aihole
}
    
func (c *伐达比VatapiCounty) BBagalkot巴加尔科特() feud.Barony {
	return c.巴加尔科特Bagalkot
}
    
func (c *伐达比VatapiCounty) BDharwar塔尔瓦尔() feud.Barony {
	return c.塔尔瓦尔Dharwar
}
    
func (c *伐达比VatapiCounty) BHalasi诃罗斯() feud.Barony {
	return c.诃罗斯Halasi
}
    
func (c *伐达比VatapiCounty) BParasgad帕拉斯迦得() feud.Barony {
	return c.帕拉斯迦得Parasgad
}
    
func (c *伐达比VatapiCounty) BVatapi伐达比() feud.Barony {
	return c.伐达比Vatapi
}
    
func (c *伐达比VatapiCounty) BVenugrama鞞纽伽罗摩() feud.Barony {
	return c.鞞纽伽罗摩Venugrama
}
    
var CVatapi伐达比 VatapiCounty = &伐达比VatapiCounty{}

func init() {
	f := CVatapi伐达比.(*伐达比VatapiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1198",
		Title:     "vatapi",
		TitleName: "伐达比",
		TitleCode: "c_vatapi",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿赫勒Aihole = BAihole阿赫勒
	f.阿赫勒Aihole.SetParent(f)
	
	f.巴加尔科特Bagalkot = BBagalkot巴加尔科特
	f.巴加尔科特Bagalkot.SetParent(f)
	
	f.塔尔瓦尔Dharwar = BDharwar塔尔瓦尔
	f.塔尔瓦尔Dharwar.SetParent(f)
	
	f.诃罗斯Halasi = BHalasi诃罗斯
	f.诃罗斯Halasi.SetParent(f)
	
	f.帕拉斯迦得Parasgad = BParasgad帕拉斯迦得
	f.帕拉斯迦得Parasgad.SetParent(f)
	
	f.伐达比Vatapi = BVatapi伐达比
	f.伐达比Vatapi.SetParent(f)
	
	f.鞞纽伽罗摩Venugrama = BVenugrama鞞纽伽罗摩
	f.鞞纽伽罗摩Venugrama.SetParent(f)
	
}
