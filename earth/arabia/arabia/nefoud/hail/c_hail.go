package hail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HailCounty interface {
    feud.County
    BAlghazalah盖扎莱() 	feud.Barony
    BAsshinan阿斯什纳() 	feud.Barony
    BBaqa拜格阿() 	feud.Barony
    BHail哈伊勒() 	feud.Barony
    BIqdah拉奇丹() 	feud.Barony
    BMawqaq毛盖格() 	feud.Barony
    BMurayfiq穆赖菲格() 	feud.Barony
    BSaban萨布安() 	feud.Barony
}

type 哈伊勒HailCounty struct {
	feud.BaseCounty
	盖扎莱Alghazalah 	feud.Barony
	阿斯什纳Asshinan 	feud.Barony
	拜格阿Baqa 	feud.Barony
	哈伊勒Hail 	feud.Barony
	拉奇丹Iqdah 	feud.Barony
	毛盖格Mawqaq 	feud.Barony
	穆赖菲格Murayfiq 	feud.Barony
	萨布安Saban 	feud.Barony
}

func (c *哈伊勒HailCounty) BAlghazalah盖扎莱() feud.Barony {
	return c.盖扎莱Alghazalah
}
    
func (c *哈伊勒HailCounty) BAsshinan阿斯什纳() feud.Barony {
	return c.阿斯什纳Asshinan
}
    
func (c *哈伊勒HailCounty) BBaqa拜格阿() feud.Barony {
	return c.拜格阿Baqa
}
    
func (c *哈伊勒HailCounty) BHail哈伊勒() feud.Barony {
	return c.哈伊勒Hail
}
    
func (c *哈伊勒HailCounty) BIqdah拉奇丹() feud.Barony {
	return c.拉奇丹Iqdah
}
    
func (c *哈伊勒HailCounty) BMawqaq毛盖格() feud.Barony {
	return c.毛盖格Mawqaq
}
    
func (c *哈伊勒HailCounty) BMurayfiq穆赖菲格() feud.Barony {
	return c.穆赖菲格Murayfiq
}
    
func (c *哈伊勒HailCounty) BSaban萨布安() feud.Barony {
	return c.萨布安Saban
}
    
var CHail哈伊勒 HailCounty = &哈伊勒HailCounty{}

func init() {
	f := CHail哈伊勒.(*哈伊勒HailCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "864",
		Title:     "hail",
		TitleName: "哈伊勒",
		TitleCode: "c_hail",
		Baronies:  map[string]feud.Barony{},
	}

	f.盖扎莱Alghazalah = BAlghazalah盖扎莱
	f.盖扎莱Alghazalah.SetParent(f)
	
	f.阿斯什纳Asshinan = BAsshinan阿斯什纳
	f.阿斯什纳Asshinan.SetParent(f)
	
	f.拜格阿Baqa = BBaqa拜格阿
	f.拜格阿Baqa.SetParent(f)
	
	f.哈伊勒Hail = BHail哈伊勒
	f.哈伊勒Hail.SetParent(f)
	
	f.拉奇丹Iqdah = BIqdah拉奇丹
	f.拉奇丹Iqdah.SetParent(f)
	
	f.毛盖格Mawqaq = BMawqaq毛盖格
	f.毛盖格Mawqaq.SetParent(f)
	
	f.穆赖菲格Murayfiq = BMurayfiq穆赖菲格
	f.穆赖菲格Murayfiq.SetParent(f)
	
	f.萨布安Saban = BSaban萨布安
	f.萨布安Saban.SetParent(f)
	
}
