package albany

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/albany/atholl"
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/albany/fife"
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/albany/gowrie"
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/albany/strathearn"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlbanyDuke interface {
    feud.Duke
    CAtholl阿瑟尔() 	atholl.AthollCounty
    CFife法夫() 	fife.FifeCounty
    CGowrie高里() 	gowrie.GowrieCounty
    CStrathearn斯特拉森() 	strathearn.StrathearnCounty
}

type 奥尔巴尼AlbanyDuke struct {
	feud.BaseDuke
	阿瑟尔Atholl 	atholl.AthollCounty
	法夫Fife 	fife.FifeCounty
	高里Gowrie 	gowrie.GowrieCounty
	斯特拉森Strathearn 	strathearn.StrathearnCounty
}

func (d *奥尔巴尼AlbanyDuke) CAtholl阿瑟尔() atholl.AthollCounty {
	return d.阿瑟尔Atholl
}
    
func (d *奥尔巴尼AlbanyDuke) CFife法夫() fife.FifeCounty {
	return d.法夫Fife
}
    
func (d *奥尔巴尼AlbanyDuke) CGowrie高里() gowrie.GowrieCounty {
	return d.高里Gowrie
}
    
func (d *奥尔巴尼AlbanyDuke) CStrathearn斯特拉森() strathearn.StrathearnCounty {
	return d.斯特拉森Strathearn
}
    
var DAlbany奥尔巴尼 AlbanyDuke = &奥尔巴尼AlbanyDuke{}

func init() {
	f := DAlbany奥尔巴尼.(*奥尔巴尼AlbanyDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "albany",
		TitleName: "奥尔巴尼",
		TitleCode: "d_albany",
		Counties:  map[string]feud.County{},
	}

	f.阿瑟尔Atholl = atholl.CAtholl阿瑟尔
	f.阿瑟尔Atholl.SetParent(f)
	
	f.法夫Fife = fife.CFife法夫
	f.法夫Fife.SetParent(f)
	
	f.高里Gowrie = gowrie.CGowrie高里
	f.高里Gowrie.SetParent(f)
	
	f.斯特拉森Strathearn = strathearn.CStrathearn斯特拉森
	f.斯特拉森Strathearn.SetParent(f)
	
}
