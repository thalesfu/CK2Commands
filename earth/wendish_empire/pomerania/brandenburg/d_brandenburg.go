package brandenburg

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/brandenburg/brandenburg"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/brandenburg/havelberg"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BrandenburgDuke interface {
    feud.Duke
    CBrandenburg勃兰登堡() 	brandenburg.BrandenburgCounty
    CHavelberg哈弗尔贝格() 	havelberg.HavelbergCounty
}

type 勃兰登堡BrandenburgDuke struct {
	feud.BaseDuke
	勃兰登堡Brandenburg 	brandenburg.BrandenburgCounty
	哈弗尔贝格Havelberg 	havelberg.HavelbergCounty
}

func (d *勃兰登堡BrandenburgDuke) CBrandenburg勃兰登堡() brandenburg.BrandenburgCounty {
	return d.勃兰登堡Brandenburg
}
    
func (d *勃兰登堡BrandenburgDuke) CHavelberg哈弗尔贝格() havelberg.HavelbergCounty {
	return d.哈弗尔贝格Havelberg
}
    
var DBrandenburg勃兰登堡 BrandenburgDuke = &勃兰登堡BrandenburgDuke{}

func init() {
	f := DBrandenburg勃兰登堡.(*勃兰登堡BrandenburgDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "brandenburg",
		TitleName: "勃兰登堡",
		TitleCode: "d_brandenburg",
		Counties:  map[string]feud.County{},
	}

	f.勃兰登堡Brandenburg = brandenburg.CBrandenburg勃兰登堡
	f.勃兰登堡Brandenburg.SetParent(f)
	
	f.哈弗尔贝格Havelberg = havelberg.CHavelberg哈弗尔贝格
	f.哈弗尔贝格Havelberg.SetParent(f)
	
}
