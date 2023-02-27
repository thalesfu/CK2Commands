package franconia

import (
	"github.com/thalesfu/CK2Commands/earth/germany/germany/franconia/bamberg"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/franconia/leiningen"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/franconia/rottenburg"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/franconia/wurzburg"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FranconiaDuke interface {
    feud.Duke
    CBamberg班贝格() 	bamberg.BambergCounty
    CLeiningen阿沙芬堡() 	leiningen.LeiningenCounty
    CRottenburg魏恩斯贝格() 	rottenburg.RottenburgCounty
    CWurzburg维尔茨堡() 	wurzburg.WurzburgCounty
}

type 法兰克尼亚FranconiaDuke struct {
	feud.BaseDuke
	班贝格Bamberg 	bamberg.BambergCounty
	阿沙芬堡Leiningen 	leiningen.LeiningenCounty
	魏恩斯贝格Rottenburg 	rottenburg.RottenburgCounty
	维尔茨堡Wurzburg 	wurzburg.WurzburgCounty
}

func (d *法兰克尼亚FranconiaDuke) CBamberg班贝格() bamberg.BambergCounty {
	return d.班贝格Bamberg
}
    
func (d *法兰克尼亚FranconiaDuke) CLeiningen阿沙芬堡() leiningen.LeiningenCounty {
	return d.阿沙芬堡Leiningen
}
    
func (d *法兰克尼亚FranconiaDuke) CRottenburg魏恩斯贝格() rottenburg.RottenburgCounty {
	return d.魏恩斯贝格Rottenburg
}
    
func (d *法兰克尼亚FranconiaDuke) CWurzburg维尔茨堡() wurzburg.WurzburgCounty {
	return d.维尔茨堡Wurzburg
}
    
var DFranconia法兰克尼亚 FranconiaDuke = &法兰克尼亚FranconiaDuke{}

func init() {
	f := DFranconia法兰克尼亚.(*法兰克尼亚FranconiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "franconia",
		TitleName: "法兰克尼亚",
		TitleCode: "d_franconia",
		Counties:  map[string]feud.County{},
	}

	f.班贝格Bamberg = bamberg.CBamberg班贝格
	f.班贝格Bamberg.SetParent(f)
	
	f.阿沙芬堡Leiningen = leiningen.CLeiningen阿沙芬堡
	f.阿沙芬堡Leiningen.SetParent(f)
	
	f.魏恩斯贝格Rottenburg = rottenburg.CRottenburg魏恩斯贝格
	f.魏恩斯贝格Rottenburg.SetParent(f)
	
	f.维尔茨堡Wurzburg = wurzburg.CWurzburg维尔茨堡
	f.维尔茨堡Wurzburg.SetParent(f)
	
}
