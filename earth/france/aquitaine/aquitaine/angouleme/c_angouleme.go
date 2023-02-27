package angouleme

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AngoulemeCounty interface {
    feud.County
    BAngouleme昂古莱姆() 	feud.Barony
    BBassac巴萨克() 	feud.Barony
    BCognac干邑() 	feud.Barony
    BFontdouce丰杜斯() 	feud.Barony
    BJarnac雅尔纳克() 	feud.Barony
    BLarochefoucauld拉罗什富科() 	feud.Barony
    BLatranchade拉特朗沙德() 	feud.Barony
    BRichemont里什蒙() 	feud.Barony
}

type 昂古莱姆AngoulemeCounty struct {
	feud.BaseCounty
	昂古莱姆Angouleme 	feud.Barony
	巴萨克Bassac 	feud.Barony
	干邑Cognac 	feud.Barony
	丰杜斯Fontdouce 	feud.Barony
	雅尔纳克Jarnac 	feud.Barony
	拉罗什富科Larochefoucauld 	feud.Barony
	拉特朗沙德Latranchade 	feud.Barony
	里什蒙Richemont 	feud.Barony
}

func (c *昂古莱姆AngoulemeCounty) BAngouleme昂古莱姆() feud.Barony {
	return c.昂古莱姆Angouleme
}
    
func (c *昂古莱姆AngoulemeCounty) BBassac巴萨克() feud.Barony {
	return c.巴萨克Bassac
}
    
func (c *昂古莱姆AngoulemeCounty) BCognac干邑() feud.Barony {
	return c.干邑Cognac
}
    
func (c *昂古莱姆AngoulemeCounty) BFontdouce丰杜斯() feud.Barony {
	return c.丰杜斯Fontdouce
}
    
func (c *昂古莱姆AngoulemeCounty) BJarnac雅尔纳克() feud.Barony {
	return c.雅尔纳克Jarnac
}
    
func (c *昂古莱姆AngoulemeCounty) BLarochefoucauld拉罗什富科() feud.Barony {
	return c.拉罗什富科Larochefoucauld
}
    
func (c *昂古莱姆AngoulemeCounty) BLatranchade拉特朗沙德() feud.Barony {
	return c.拉特朗沙德Latranchade
}
    
func (c *昂古莱姆AngoulemeCounty) BRichemont里什蒙() feud.Barony {
	return c.里什蒙Richemont
}
    
var CAngouleme昂古莱姆 AngoulemeCounty = &昂古莱姆AngoulemeCounty{}

func init() {
	f := CAngouleme昂古莱姆.(*昂古莱姆AngoulemeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "148",
		Title:     "angouleme",
		TitleName: "昂古莱姆",
		TitleCode: "c_angouleme",
		Baronies:  map[string]feud.Barony{},
	}

	f.昂古莱姆Angouleme = BAngouleme昂古莱姆
	f.昂古莱姆Angouleme.SetParent(f)
	
	f.巴萨克Bassac = BBassac巴萨克
	f.巴萨克Bassac.SetParent(f)
	
	f.干邑Cognac = BCognac干邑
	f.干邑Cognac.SetParent(f)
	
	f.丰杜斯Fontdouce = BFontdouce丰杜斯
	f.丰杜斯Fontdouce.SetParent(f)
	
	f.雅尔纳克Jarnac = BJarnac雅尔纳克
	f.雅尔纳克Jarnac.SetParent(f)
	
	f.拉罗什富科Larochefoucauld = BLarochefoucauld拉罗什富科
	f.拉罗什富科Larochefoucauld.SetParent(f)
	
	f.拉特朗沙德Latranchade = BLatranchade拉特朗沙德
	f.拉特朗沙德Latranchade.SetParent(f)
	
	f.里什蒙Richemont = BRichemont里什蒙
	f.里什蒙Richemont.SetParent(f)
	
}
