package ustug

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/ustug/bolshaya"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/ustug/syktyvkar"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/ustug/veliky_ustug"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/ustug/zyriane"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UstugDuke interface {
    feud.Duke
    CBolshaya维列季() 	bolshaya.BolshayaCounty
    CSyktyvkar瑟克特夫卡尔() 	syktyvkar.SyktyvkarCounty
    CVeliky_ustug大乌斯秋格() 	veliky_ustug.Veliky_ustugCounty
    CZyriane伊伦斯克() 	zyriane.ZyrianeCounty
}

type 大乌斯秋格UstugDuke struct {
	feud.BaseDuke
	维列季Bolshaya 	bolshaya.BolshayaCounty
	瑟克特夫卡尔Syktyvkar 	syktyvkar.SyktyvkarCounty
	大乌斯秋格Veliky_ustug 	veliky_ustug.Veliky_ustugCounty
	伊伦斯克Zyriane 	zyriane.ZyrianeCounty
}

func (d *大乌斯秋格UstugDuke) CBolshaya维列季() bolshaya.BolshayaCounty {
	return d.维列季Bolshaya
}
    
func (d *大乌斯秋格UstugDuke) CSyktyvkar瑟克特夫卡尔() syktyvkar.SyktyvkarCounty {
	return d.瑟克特夫卡尔Syktyvkar
}
    
func (d *大乌斯秋格UstugDuke) CVeliky_ustug大乌斯秋格() veliky_ustug.Veliky_ustugCounty {
	return d.大乌斯秋格Veliky_ustug
}
    
func (d *大乌斯秋格UstugDuke) CZyriane伊伦斯克() zyriane.ZyrianeCounty {
	return d.伊伦斯克Zyriane
}
    
var DUstug大乌斯秋格 UstugDuke = &大乌斯秋格UstugDuke{}

func init() {
	f := DUstug大乌斯秋格.(*大乌斯秋格UstugDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ustug",
		TitleName: "大乌斯秋格",
		TitleCode: "d_ustug",
		Counties:  map[string]feud.County{},
	}

	f.维列季Bolshaya = bolshaya.CBolshaya维列季
	f.维列季Bolshaya.SetParent(f)
	
	f.瑟克特夫卡尔Syktyvkar = syktyvkar.CSyktyvkar瑟克特夫卡尔
	f.瑟克特夫卡尔Syktyvkar.SetParent(f)
	
	f.大乌斯秋格Veliky_ustug = veliky_ustug.CVeliky_ustug大乌斯秋格
	f.大乌斯秋格Veliky_ustug.SetParent(f)
	
	f.伊伦斯克Zyriane = zyriane.CZyriane伊伦斯克
	f.伊伦斯克Zyriane.SetParent(f)
	
}
