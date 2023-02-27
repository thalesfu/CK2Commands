package kuyavia

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/kuyavia/dobrzynska"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/kuyavia/kujawy"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/kuyavia/sieradzka"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/kuyavia/wielunska"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KuyaviaDuke interface {
    feud.Duke
    CDobrzynska多布任() 	dobrzynska.DobrzynskaCounty
    CKujawy库亚维() 	kujawy.KujawyCounty
    CSieradzka谢拉兹() 	sieradzka.SieradzkaCounty
    CWielunska维隆() 	wielunska.WielunskaCounty
}

type 库亚维亚KuyaviaDuke struct {
	feud.BaseDuke
	多布任Dobrzynska 	dobrzynska.DobrzynskaCounty
	库亚维Kujawy 	kujawy.KujawyCounty
	谢拉兹Sieradzka 	sieradzka.SieradzkaCounty
	维隆Wielunska 	wielunska.WielunskaCounty
}

func (d *库亚维亚KuyaviaDuke) CDobrzynska多布任() dobrzynska.DobrzynskaCounty {
	return d.多布任Dobrzynska
}
    
func (d *库亚维亚KuyaviaDuke) CKujawy库亚维() kujawy.KujawyCounty {
	return d.库亚维Kujawy
}
    
func (d *库亚维亚KuyaviaDuke) CSieradzka谢拉兹() sieradzka.SieradzkaCounty {
	return d.谢拉兹Sieradzka
}
    
func (d *库亚维亚KuyaviaDuke) CWielunska维隆() wielunska.WielunskaCounty {
	return d.维隆Wielunska
}
    
var DKuyavia库亚维亚 KuyaviaDuke = &库亚维亚KuyaviaDuke{}

func init() {
	f := DKuyavia库亚维亚.(*库亚维亚KuyaviaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kuyavia",
		TitleName: "库亚维亚",
		TitleCode: "d_kuyavia",
		Counties:  map[string]feud.County{},
	}

	f.多布任Dobrzynska = dobrzynska.CDobrzynska多布任
	f.多布任Dobrzynska.SetParent(f)
	
	f.库亚维Kujawy = kujawy.CKujawy库亚维
	f.库亚维Kujawy.SetParent(f)
	
	f.谢拉兹Sieradzka = sieradzka.CSieradzka谢拉兹
	f.谢拉兹Sieradzka.SetParent(f)
	
	f.维隆Wielunska = wielunska.CWielunska维隆
	f.维隆Wielunska.SetParent(f)
	
}
