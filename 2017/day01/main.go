package main

import (
	"fmt"
	"log"
)

func Sum(digits string) int {
	digitsInt := make([]int, len(digits)+1)
	for i, d := range digits {
		switch d {
		case '1':
			digitsInt[i] = 1
		case '2':
			digitsInt[i] = 2
		case '3':
			digitsInt[i] = 3
		case '4':
			digitsInt[i] = 4
		case '5':
			digitsInt[i] = 5
		case '6':
			digitsInt[i] = 6
		case '7':
			digitsInt[i] = 7
		case '8':
			digitsInt[i] = 8
		case '9':
			digitsInt[i] = 9
		case '0':
			digitsInt[i] = 0
		}
	}
	digitsInt[len(digitsInt)-1] = digitsInt[0]

	sum := 0
	for i := 0; i < len(digitsInt)-1; i++ {
		if digitsInt[i] == digitsInt[i+1] {
			sum += digitsInt[i]
		}
	}

	return sum
}

func Sum2(digits string) int {
	digitsInt := make([]int, len(digits))
	digitsIntR := make([]int, len(digits))
	for i, d := range digits {
		switch d {
		case '1':
			digitsInt[i] = 1
		case '2':
			digitsInt[i] = 2
		case '3':
			digitsInt[i] = 3
		case '4':
			digitsInt[i] = 4
		case '5':
			digitsInt[i] = 5
		case '6':
			digitsInt[i] = 6
		case '7':
			digitsInt[i] = 7
		case '8':
			digitsInt[i] = 8
		case '9':
			digitsInt[i] = 9
		case '0':
			digitsInt[i] = 0
		}
	}
	half := len(digitsInt) / 2
	for i := range digitsInt {
		if i < half {
			digitsIntR[i] = digitsInt[i+half]
		} else {
			digitsIntR[i] = digitsInt[i-half]
		}
	}

	sum := 0
	for i := 0; i < len(digitsInt); i++ {
		if digitsInt[i] == digitsIntR[i] {
			sum += digitsInt[i]
		}
	}

	return sum
}
func main() {
	t1 := Sum("1122")
	if t1 != 3 {
		log.Fatal("Didn't sum 1122", t1)
	}
	t2 := Sum("1111")
	if t2 != 4 {
		log.Fatal("Didn't sum 1111", t1)
	}
	t3 := Sum("1234")
	if t3 != 0 {
		log.Fatal("Didn't sum 1234", t1)
	}
	t4 := Sum("91212129")
	if t4 != 9 {
		log.Fatal("Didn't sum 91212129", t1)
	}

	fmt.Println("SUM", Sum("57276274387944537823652626177853384411146325384494935924454336611953119173638191671326254832624841593421667683474349154668177743437745965461678636631863541462893547616877914914662358836365421198516263335926544716331814125295712581158399321372683742773423626286669759415959391374744214595682795818615532673877868424196926497731144319736445141728123322962547288572434564178492753681842244888368542423832228211172842456231275738182764232265933625119312598161192193214898949267765417468348935134618964683127194391796165368145548814473129857697989322621368744725685183346825333247866734735894493395218781464346951777873929898961358796274889826894529599645442657423438562423853247543621565468819799931598754753467593832328147439341586125262733737128386961596394728159719292787597426898945198788211417854662948358422729471312456437778978749753927251431677533575752312447488337156956217451965643454445329758327129966657189332824969141448538681979632611199385896965946849725421978137753366252459914913637858783146735469758716752765718189175583956476935185985918536318424248425426398158278111751711911227818826766177996223718837428972784328925743869885232266127727865267881592395643836999244218345184474613129823933659422223685422732186536199153988717455568523781673393698356967355875123554797755491181791593156433735591529495984256519631187849654633243225118132152549712643273819314433877592644693826861523243946998615722951182474773173215527598949553185313259992227879964482121769617218685394776778423378182462422788277997523913176326468957342296368178321958626168785578977414537368686438348124283789748775163821457641135163495649331144436157836647912852483177542224864952271874645274572426458614384917923623627532487625396914111582754953944965462576624728896917137599778828769958626788685374749661741223741834844643725486925886933118382649581481351844943368484853956759877215252766294896496444835264357169642341291412768946589781812493421379575569593678354241223363739129813633236996588711791919421574583924743119867622229659211793468744163297478952475933163259769578345894367855534294493613767564497137369969315192443795512585"))

	t1 = Sum2("1212")
	if t1 != 6 {
		log.Fatal("Didn't sum 1212:", t1)
	}
	t2 = Sum2("1221")
	if t2 != 0 {
		log.Fatal("Didn't sum 1111:", t1)
	}
	t3 = Sum2("123425")
	if t3 != 4 {
		log.Fatal("Didn't sum 123425:", t1)
	}
	t4 = Sum2("123123")
	if t4 != 12 {
		log.Fatal("Didn't sum 123123:", t1)
	}
	t5 := Sum2("12131415")
	if t5 != 4 {
		log.Fatal("Didn't sum 12131415:", t1)
	}
	fmt.Println("SUM2", Sum2("57276274387944537823652626177853384411146325384494935924454336611953119173638191671326254832624841593421667683474349154668177743437745965461678636631863541462893547616877914914662358836365421198516263335926544716331814125295712581158399321372683742773423626286669759415959391374744214595682795818615532673877868424196926497731144319736445141728123322962547288572434564178492753681842244888368542423832228211172842456231275738182764232265933625119312598161192193214898949267765417468348935134618964683127194391796165368145548814473129857697989322621368744725685183346825333247866734735894493395218781464346951777873929898961358796274889826894529599645442657423438562423853247543621565468819799931598754753467593832328147439341586125262733737128386961596394728159719292787597426898945198788211417854662948358422729471312456437778978749753927251431677533575752312447488337156956217451965643454445329758327129966657189332824969141448538681979632611199385896965946849725421978137753366252459914913637858783146735469758716752765718189175583956476935185985918536318424248425426398158278111751711911227818826766177996223718837428972784328925743869885232266127727865267881592395643836999244218345184474613129823933659422223685422732186536199153988717455568523781673393698356967355875123554797755491181791593156433735591529495984256519631187849654633243225118132152549712643273819314433877592644693826861523243946998615722951182474773173215527598949553185313259992227879964482121769617218685394776778423378182462422788277997523913176326468957342296368178321958626168785578977414537368686438348124283789748775163821457641135163495649331144436157836647912852483177542224864952271874645274572426458614384917923623627532487625396914111582754953944965462576624728896917137599778828769958626788685374749661741223741834844643725486925886933118382649581481351844943368484853956759877215252766294896496444835264357169642341291412768946589781812493421379575569593678354241223363739129813633236996588711791919421574583924743119867622229659211793468744163297478952475933163259769578345894367855534294493613767564497137369969315192443795512585"))
}
