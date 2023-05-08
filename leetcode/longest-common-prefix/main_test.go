package main

import "testing"

var testcase = [2]string{
	"abcdefghijklmnopqrstuvwxyz3354254e656453442432",
	"abcdefghijklmnopqrstuvwxyz8f85f40f08w3n80dfw48fnufsof",
}

var longtestcases = []string{
	"abcdefghijklmnopqrstuvwxyz12504851976145641821776324682545226035664858215592569028",
	"abcdefghijklmnopqrstuvwxyz305679473380962920948337137091387347182624111440481981425383",
	"abcdefghijklmnopqrstuvwxyz67429379302440975119779893977156961121582191532223501315",
	"abcdefghijklmnopqrstuvwxyz798923579807845562591757733287585836793617307892992345",
	"abcdefghijklmnopqrstuvwxyz88291609339663365240109224560452111175394521799939619364",
	"abcdefghijklmnopqrstuvwxyz1318655727861453902661635448144427433913207034181241474",
	"abcdefghijklmnopqrstuvwxyz221267417587589694562180102044938872264888027697645156121",
	"abcdefghijklmnopqrstuvwxyz59253730625367956844188772974746958510021433329473577924744",
	"abcdefghijklmnopqrstuvwxyz96941767384726681716477059298516937567836137038984090552827",
	"abcdefghijklmnopqrstuvwxyz88711098981282376431403867656106545501766667022101326752388",
	"abcdefghijklmnopqrstuvwxyz8068547759845975967586214854338353732120801432728291785130",
	"abcdefghijklmnopqrstuvwxyz9845308027024106620235212070839662265810028287573730717179",
	"abcdefghijklmnopqrstuvwxyz809313682066892337478039423251528135119525929833627411895126",
	"abcdefghijklmnopqrstuvwxyz845165425228525360942342719648922263243376679826978593395",
	"abcdefghijklmnopqrstuvwxyz6335685391018905216737094223028885011585768240958696797061",
	"abcdefghijklmnopqrstuvwxyz63948610105142828528051312388524406235158050394053921144",
	"abcdefghijklmnopqrstuvwxyz2898992774354580353535267564598527897243489725216514254606",
	"abcdefghijklmnopqrstuvwxyz82275186283549457560324575655310027992610058423266103411467413",
	"abcdefghijklmnopqrstuvwxyz443010873298501910645891456462128050824852467255274910026156",
	"abcdefghijklmnopqrstuvwxyz9276831648844010090978224975143449781949884954094735559998",
	"abcdefghijklmnopqrstuvwxyz93341463381243659593505999199050197323464468744546424534112",
	"abcdefghijklmnopqrstuvwxyz87129541413187171479623230597681667851664603486611095873",
	"abcdefghijklmnopqrstuvwxyz9294059753246182407175887616173366744622293469462918382",
	"abcdefghijklmnopqrstuvwxyz16265128844211711161007594982811151698944274870133542388785",
	"abcdefghijklmnopqrstuvwxyz3492378318026158861281536579124139841983392934546177944862",
	"abcdefghijklmnopqrstuvwxyz90190823176330964158226480984464280135680339653128785620",
	"abcdefghijklmnopqrstuvwxyz225249102922317449836683823458328754201910099266752127053864",
	"abcdefghijklmnopqrstuvwxyz1633137113152434341269805030539523838944813768669173136",
	"abcdefghijklmnopqrstuvwxyz67155518752549226773894392354196789382355832437887918789",
	"abcdefghijklmnopqrstuvwxyz75282962476080345701452193047105869100446537723044245103141",
	"abcdefghijklmnopqrstuvwxyz20684286460968553126693415155927837442933176054952982978",
	"abcdefghijklmnopqrstuvwxyz845174484394776956112124166849910033374147301139327696588952",
	"abcdefghijklmnopqrstuvwxyz9443828514683445434504369873166585068503916279810010014461432",
	"abcdefghijklmnopqrstuvwxyz84845024649328366161762931908947471868208858952644661239425",
	"abcdefghijklmnopqrstuvwxyz846197383095523953161453815952729651315185823100215399674030",
	"abcdefghijklmnopqrstuvwxyz451637297654806965534789694080157531823332646976944772",
	"abcdefghijklmnopqrstuvwxyz7570827231133592272617346515926172321678115332189383822051",
	"abcdefghijklmnopqrstuvwxyz70868165546227318381775978851522629345537341592423356116934",
	"abcdefghijklmnopqrstuvwxyz6866417391024743856483857519801981701026437481005339543477",
	"abcdefghijklmnopqrstuvwxyz82557134739244756096736338588945297717302114391650552250736",
	"abcdefghijklmnopqrstuvwxyz16464391283642116993898010548353147698025206552981831267",
	"abcdefghijklmnopqrstuvwxyz5090655659180279333341740575889662161037959011307664595364",
	"abcdefghijklmnopqrstuvwxyz8334837613616944917313665896902944779100254773992050918382",
	"abcdefghijklmnopqrstuvwxyz75136781716188572019164871100866151218124447227784684373",
	"abcdefghijklmnopqrstuvwxyz798015564125694468067907941896289050855876887338220925172",
	"abcdefghijklmnopqrstuvwxyz783366153863959742351924496861154773796863060539173587723",
	"abcdefghijklmnopqrstuvwxyz221232763536966302555406380527782546971566795240858688624",
	"abcdefghijklmnopqrstuvwxyz439669301144394430395428232811718841984155131937513793",
	"abcdefghijklmnopqrstuvwxyz209679315817290238161392819738222140479286728189172779026",
	"abcdefghijklmnopqrstuvwxyz9754359334152370512765854452892719756121726552826724936378",
	"abcdefghijklmnopqrstuvwxyz312915568296151366239358379225489822683377427549386213",
	"abcdefghijklmnopqrstuvwxyz7265243483978867419690761635742725344116176991902424432474",
	"abcdefghijklmnopqrstuvwxyz968075539942954841533570517851821275509554528616678634021",
	"abcdefghijklmnopqrstuvwxyz7263978287363879254414788837711546348775792519312463445210",
	"abcdefghijklmnopqrstuvwxyz4153871866198292746752133823593292992312229180804582211",
	"abcdefghijklmnopqrstuvwxyz26406222936645461751001098674453756473031192844422597243",
	"abcdefghijklmnopqrstuvwxyz3912769055333831133476012363861874797876988530628034181713",
	"abcdefghijklmnopqrstuvwxyz9613108840669752489799299813242678482794676983384081494242",
	"abcdefghijklmnopqrstuvwxyz217379974994939872742503532968686468342735995979897210",
	"abcdefghijklmnopqrstuvwxyz7243019368686789593786988275602784553569143166811686432",
	"abcdefghijklmnopqrstuvwxyz76294688871627908791465750663055933089978786261726380912718",
	"abcdefghijklmnopqrstuvwxyz673618779127183947774616989185064174471273744595699191007171",
	"abcdefghijklmnopqrstuvwxyz1110048137863836553235823399121100793853889761450169401986",
	"abcdefghijklmnopqrstuvwxyz111044546969551106738742427501710062374045789664769149697653",
	"abcdefghijklmnopqrstuvwxyz2623103297436089211411844027359862354115517717110065836871",
	"abcdefghijklmnopqrstuvwxyz36815628342680491354726085138100661045311412319178245916494",
	"abcdefghijklmnopqrstuvwxyz26198546171697184037029307325746529339250486089942894271072",
	"abcdefghijklmnopqrstuvwxyz178343140289993416383677933544841122475745149693339527590",
	"abcdefghijklmnopqrstuvwxyz8413688244743969922547321932598253293100359896080244645516",
	"abcdefghijklmnopqrstuvwxyz9737839570457686376069131283767279496211893238376445453681",
	"abcdefghijklmnopqrstuvwxyz779196151765892956335674361185174840288252725067199472471",
	"abcdefghijklmnopqrstuvwxyz1568649235321838602599688288071227262697599444793273615",
	"abcdefghijklmnopqrstuvwxyz6333834723175212319166825638729686153083452131593874182",
	"abcdefghijklmnopqrstuvwxyz696610079350282765924097927373164567534186645318313694731",
	"abcdefghijklmnopqrstuvwxyz6941380273675468274910246283421383382029301369626594364023",
	"abcdefghijklmnopqrstuvwxyz22895519758031713703250802262316776189585828672510062613298",
	"abcdefghijklmnopqrstuvwxyz3312478224987576197864798298687384163281735436367847251053",
	"abcdefghijklmnopqrstuvwxyz895129885862653030981359310687046132327382553090422651014",
	"abcdefghijklmnopqrstuvwxyz53567450868114713687134533100351078478279814941737084837787",
	"abcdefghijklmnopqrstuvwxyz88332263256026898951854390397741763451962197938217593198310",
	"abcdefghijklmnopqrstuvwxyz4896595568675388739502415308326157318559818697549492318100",
	"abcdefghijklmnopqrstuvwxyz261481542676446737862393418347387495116184602561159110155",
	"abcdefghijklmnopqrstuvwxyz8268966194815911259365733415984339567458859571988789558536",
	"abcdefghijklmnopqrstuvwxyz881368451919115463525854535863304518654627303435614324713",
	"abcdefghijklmnopqrstuvwxyz53837258661009654666728682798854182545234417755534872450391",
	"abcdefghijklmnopqrstuvwxyz801967786986743171258844388468599849168022927563712760363278",
	"abcdefghijklmnopqrstuvwxyz6555314479288861998743071675653439559799267954253965392",
	"abcdefghijklmnopqrstuvwxyz1465696368868567231422739326299843736210343013314690417013",
	"abcdefghijklmnopqrstuvwxyz79374332114758760724045516225782920422015477145569431166852",
	"abcdefghijklmnopqrstuvwxyz9279798568393061329241229146144919188601061253168571180363",
	"abcdefghijklmnopqrstuvwxyz33974432109160467139789855894097835649573636916983892847920",
	"abcdefghijklmnopqrstuvwxyz5376013586219960854225266327844539655091192287582189452744",
	"abcdefghijklmnopqrstuvwxyz8473529317402732536596383597947065669384756466999425375454",
	"abcdefghijklmnopqrstuvwxyz248468781991737825118535539344341001163991246205269244831341",
	"abcdefghijklmnopqrstuvwxyz8130233417311665929375743368666513975515311563662583576129",
	"abcdefghijklmnopqrstuvwxyz6340516573468718751811513421768567235808571718333331022",
	"abcdefghijklmnopqrstuvwxyz45522533228053735922435957312314972121728836618843729558967",
	"abcdefghijklmnopqrstuvwxyz684681797618425523327313703930165118915860239583598756639",
	"abcdefghijklmnopqrstuvwxyz8243905574942728162076962223931090269792121791472155238254",
	"abcdefghijklmnopqrstuvwxyz6995272811968759533722587244367318311003578100292719943812",
}

var out string

func BenchmarkLongestCommonPrefixTwo(b *testing.B) {
	var val string
	for n := 0; n < b.N; n++ {
		val = longestCommonPrefixTwo(testcase[0], testcase[1])
	}
	out = val
}

func BenchmarkLongestCommonPrefixTwoNewBuilder(b *testing.B) {
	var val string
	for n := 0; n < b.N; n++ {
		val = longestCommonPrefixTwoNewBuilder(testcase[0], testcase[1])
	}
	out = val
}

func BenchmarkLongestCommonPrefixTwoStringAddition(b *testing.B) {
	var val string
	for n := 0; n < b.N; n++ {
		val = longestCommonPrefixTwoStringAddition(testcase[0], testcase[1])
	}
	out = val
}

func BenchmarkLongestCommonPrefixBuilder1(b *testing.B) {
	var val string
	for n := 0; n < b.N; n++ {
		val = longestCommonPrefixB(longtestcases, longestCommonPrefixTwo)
	}
	out = val
}
func BenchmarkLongestCommonPrefixBuilderNew(b *testing.B) {
	var val string
	for n := 0; n < b.N; n++ {
		val = longestCommonPrefixB(longtestcases, longestCommonPrefixTwoNewBuilder)
	}
	out = val
}
func BenchmarkLongestCommonPrefixStringAddition(b *testing.B) {
	var val string
	for n := 0; n < b.N; n++ {
		val = longestCommonPrefixB(longtestcases, longestCommonPrefixTwoStringAddition)
	}
	out = val
}
