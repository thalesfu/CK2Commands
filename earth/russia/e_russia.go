package russia

import (
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov"
	"github.com/thalesfu/CK2Commands/earth/russia/galicia_volhynia"
	"github.com/thalesfu/CK2Commands/earth/russia/rus"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RussiaEmpire interface {
    feud.Empire
    KChernigov切尔尼戈夫() 	chernigov.ChernigovKingdom
    KGalicia_volhynia加利西亚_沃里尼亚() 	galicia_volhynia.Galicia_volhyniaKingdom
    KRus诺夫哥罗德() 	rus.RusKingdom
    KRuthenia鲁塞尼亚() 	ruthenia.RutheniaKingdom
    KVladimir弗拉基米尔() 	vladimir.VladimirKingdom
}

type 罗斯帝国RussiaEmpire struct {
	feud.BaseEmpire
	切尔尼戈夫Chernigov 	chernigov.ChernigovKingdom
	加利西亚_沃里尼亚Galicia_volhynia 	galicia_volhynia.Galicia_volhyniaKingdom
	诺夫哥罗德Rus 	rus.RusKingdom
	鲁塞尼亚Ruthenia 	ruthenia.RutheniaKingdom
	弗拉基米尔Vladimir 	vladimir.VladimirKingdom
}

func (e *罗斯帝国RussiaEmpire) KChernigov切尔尼戈夫() chernigov.ChernigovKingdom {
	return e.切尔尼戈夫Chernigov
}
    
func (e *罗斯帝国RussiaEmpire) KGalicia_volhynia加利西亚_沃里尼亚() galicia_volhynia.Galicia_volhyniaKingdom {
	return e.加利西亚_沃里尼亚Galicia_volhynia
}
    
func (e *罗斯帝国RussiaEmpire) KRus诺夫哥罗德() rus.RusKingdom {
	return e.诺夫哥罗德Rus
}
    
func (e *罗斯帝国RussiaEmpire) KRuthenia鲁塞尼亚() ruthenia.RutheniaKingdom {
	return e.鲁塞尼亚Ruthenia
}
    
func (e *罗斯帝国RussiaEmpire) KVladimir弗拉基米尔() vladimir.VladimirKingdom {
	return e.弗拉基米尔Vladimir
}
    
var ERussia罗斯帝国 RussiaEmpire = &罗斯帝国RussiaEmpire{}

func init() {
	f := ERussia罗斯帝国.(*罗斯帝国RussiaEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "russia",
		TitleName: "罗斯帝国",
		TitleCode: "e_russia",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.切尔尼戈夫Chernigov = chernigov.KChernigov切尔尼戈夫
	f.切尔尼戈夫Chernigov.SetParent(f)
	f.加利西亚_沃里尼亚Galicia_volhynia = galicia_volhynia.KGalicia_volhynia加利西亚_沃里尼亚
	f.加利西亚_沃里尼亚Galicia_volhynia.SetParent(f)
	f.诺夫哥罗德Rus = rus.KRus诺夫哥罗德
	f.诺夫哥罗德Rus.SetParent(f)
	f.鲁塞尼亚Ruthenia = ruthenia.KRuthenia鲁塞尼亚
	f.鲁塞尼亚Ruthenia.SetParent(f)
	f.弗拉基米尔Vladimir = vladimir.KVladimir弗拉基米尔
	f.弗拉基米尔Vladimir.SetParent(f)
}
