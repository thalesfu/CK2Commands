package uppland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UpplandCounty interface {
    feud.County
    BAlsno阿尔斯内() 	feud.Barony
    BBirka比尔卡() 	feud.Barony
    BEnkoping恩雪平() 	feud.Barony
    BHatuna霍图纳() 	feud.Barony
    BOsteras厄斯特罗斯() 	feud.Barony
    BSigtuna锡格蒂纳() 	feud.Barony
    BStockholm斯德哥尔摩() 	feud.Barony
    BUppsala乌普萨拉() 	feud.Barony
    BVaksala瓦克萨拉() 	feud.Barony
}

type 乌普兰UpplandCounty struct {
	feud.BaseCounty
	阿尔斯内Alsno 	feud.Barony
	比尔卡Birka 	feud.Barony
	恩雪平Enkoping 	feud.Barony
	霍图纳Hatuna 	feud.Barony
	厄斯特罗斯Osteras 	feud.Barony
	锡格蒂纳Sigtuna 	feud.Barony
	斯德哥尔摩Stockholm 	feud.Barony
	乌普萨拉Uppsala 	feud.Barony
	瓦克萨拉Vaksala 	feud.Barony
}

func (c *乌普兰UpplandCounty) BAlsno阿尔斯内() feud.Barony {
	return c.阿尔斯内Alsno
}
    
func (c *乌普兰UpplandCounty) BBirka比尔卡() feud.Barony {
	return c.比尔卡Birka
}
    
func (c *乌普兰UpplandCounty) BEnkoping恩雪平() feud.Barony {
	return c.恩雪平Enkoping
}
    
func (c *乌普兰UpplandCounty) BHatuna霍图纳() feud.Barony {
	return c.霍图纳Hatuna
}
    
func (c *乌普兰UpplandCounty) BOsteras厄斯特罗斯() feud.Barony {
	return c.厄斯特罗斯Osteras
}
    
func (c *乌普兰UpplandCounty) BSigtuna锡格蒂纳() feud.Barony {
	return c.锡格蒂纳Sigtuna
}
    
func (c *乌普兰UpplandCounty) BStockholm斯德哥尔摩() feud.Barony {
	return c.斯德哥尔摩Stockholm
}
    
func (c *乌普兰UpplandCounty) BUppsala乌普萨拉() feud.Barony {
	return c.乌普萨拉Uppsala
}
    
func (c *乌普兰UpplandCounty) BVaksala瓦克萨拉() feud.Barony {
	return c.瓦克萨拉Vaksala
}
    
var CUppland乌普兰 UpplandCounty = &乌普兰UpplandCounty{}

func init() {
	f := CUppland乌普兰.(*乌普兰UpplandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "290",
		Title:     "uppland",
		TitleName: "乌普兰",
		TitleCode: "c_uppland",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔斯内Alsno = BAlsno阿尔斯内
	f.阿尔斯内Alsno.SetParent(f)
	
	f.比尔卡Birka = BBirka比尔卡
	f.比尔卡Birka.SetParent(f)
	
	f.恩雪平Enkoping = BEnkoping恩雪平
	f.恩雪平Enkoping.SetParent(f)
	
	f.霍图纳Hatuna = BHatuna霍图纳
	f.霍图纳Hatuna.SetParent(f)
	
	f.厄斯特罗斯Osteras = BOsteras厄斯特罗斯
	f.厄斯特罗斯Osteras.SetParent(f)
	
	f.锡格蒂纳Sigtuna = BSigtuna锡格蒂纳
	f.锡格蒂纳Sigtuna.SetParent(f)
	
	f.斯德哥尔摩Stockholm = BStockholm斯德哥尔摩
	f.斯德哥尔摩Stockholm.SetParent(f)
	
	f.乌普萨拉Uppsala = BUppsala乌普萨拉
	f.乌普萨拉Uppsala.SetParent(f)
	
	f.瓦克萨拉Vaksala = BVaksala瓦克萨拉
	f.瓦克萨拉Vaksala.SetParent(f)
	
}
