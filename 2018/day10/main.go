package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	Do(demo, "demo1.png")
	Do(input, "input1.png")
}

var lineRegex = regexp.MustCompile(`position=< ?(-?\d+), *(-?\d+)> velocity=< *(-?\d+), *(-?\d+)>`)

func Do(input, name string) {
	points := []*Point{}
	for _, line := range strings.Split(input, "\n") {
		parts := lineRegex.FindStringSubmatch(line)[1:]
		p := &Point{}
		p.OriginX, _ = strconv.Atoi(parts[0])
		p.OriginY, _ = strconv.Atoi(parts[1])
		p.VelX, _ = strconv.Atoi(parts[2])
		p.VelY, _ = strconv.Atoi(parts[3])
		p.X = p.OriginX
		p.Y = p.OriginY
		points = append(points, p)
	}
	for i := 0; i < 25000; i++ {
		m := map[int]bool{}
		xmax, ymax := 0, 0
		for _, p := range points {
			x, y := p.Step()
			m[y] = true
			if x > xmax {
				xmax = x
			}
			if y > ymax {
				ymax = y
			}
		}
		if len(m) <= 10 {
			fmt.Println(i+1, "Seconds")
			img := image.NewRGBA(image.Rect(0, 0, xmax+10, ymax+10))
			black := color.RGBA{0, 0, 0, 255}
			for _, p := range points {
				img.Set(p.X+5, p.Y+5, black)
			}
			f, _ := os.Create(name)
			png.Encode(f, img)
			f.Close()
			return
		}
	}
}

type Point struct {
	OriginX, OriginY int
	VelX, VelY       int
	X, Y             int
}

func (p *Point) Step() (int, int) {
	p.X += p.VelX
	p.Y += p.VelY
	return p.X, p.Y
}

var demo = `position=< 9,  1> velocity=< 0,  2>
position=< 7,  0> velocity=<-1,  0>
position=< 3, -2> velocity=<-1,  1>
position=< 6, 10> velocity=<-2, -1>
position=< 2, -4> velocity=< 2,  2>
position=<-6, 10> velocity=< 2, -2>
position=< 1,  8> velocity=< 1, -1>
position=< 1,  7> velocity=< 1,  0>
position=<-3, 11> velocity=< 1, -2>
position=< 7,  6> velocity=<-1, -1>
position=<-2,  3> velocity=< 1,  0>
position=<-4,  3> velocity=< 2,  0>
position=<10, -3> velocity=<-1,  1>
position=< 5, 11> velocity=< 1, -2>
position=< 4,  7> velocity=< 0, -1>
position=< 8, -2> velocity=< 0,  1>
position=<15,  0> velocity=<-2,  0>
position=< 1,  6> velocity=< 1,  0>
position=< 8,  9> velocity=< 0, -1>
position=< 3,  3> velocity=<-1,  1>
position=< 0,  5> velocity=< 0, -1>
position=<-2,  2> velocity=< 2,  0>
position=< 5, -2> velocity=< 1,  2>
position=< 1,  4> velocity=< 2,  1>
position=<-2,  7> velocity=< 2, -2>
position=< 3,  6> velocity=<-1, -1>
position=< 5,  0> velocity=< 1,  0>
position=<-6,  0> velocity=< 2,  0>
position=< 5,  9> velocity=< 1, -2>
position=<14,  7> velocity=<-2,  0>
position=<-3,  6> velocity=< 2, -1>`

var input = `position=<-52775,  31912> velocity=< 5, -3>
position=<-52816,  10731> velocity=< 5, -1>
position=< 42573, -31652> velocity=<-4,  3>
position=<-31611,  31918> velocity=< 3, -3>
position=< 42585, -31656> velocity=<-4,  3>
position=<-10428,  53102> velocity=< 1, -5>
position=< 31958, -10460> velocity=<-3,  1>
position=<-21020,  10723> velocity=< 2, -1>
position=<-52770,  10731> velocity=< 5, -1>
position=< 21400,  10728> velocity=<-2, -1>
position=<-10427, -42253> velocity=< 1,  4>
position=< 21343,  42515> velocity=<-2, -4>
position=<-10417, -52846> velocity=< 1,  5>
position=< 42558, -42251> velocity=<-4,  4>
position=<-21008,  53106> velocity=< 2, -5>
position=<-10443,  31912> velocity=< 1, -3>
position=<-31584,  31919> velocity=< 3, -3>
position=<-31604,  10726> velocity=< 3, -1>
position=<-52786, -21055> velocity=< 5,  2>
position=<-10394,  31918> velocity=< 1, -3>
position=<-52786,  53109> velocity=< 5, -5>
position=<-52770,  10722> velocity=< 5, -1>
position=<-21007, -10464> velocity=< 2,  1>
position=< 21392, -31657> velocity=<-2,  3>
position=< 53180,  31913> velocity=<-5, -3>
position=<-10393,  21319> velocity=< 1, -2>
position=<-10430,  10724> velocity=< 1, -1>
position=<-20985, -42245> velocity=< 2,  4>
position=<-31610, -10468> velocity=< 3,  1>
position=<-20989, -42247> velocity=< 2,  4>
position=<-21006,  21317> velocity=< 2, -2>
position=< 10752, -31651> velocity=<-1,  3>
position=< 42577,  31912> velocity=<-4, -3>
position=<-42196, -21059> velocity=< 4,  2>
position=< 10744, -10460> velocity=<-1,  1>
position=<-10443, -31658> velocity=< 1,  3>
position=< 53185,  31912> velocity=<-5, -3>
position=<-52790, -21063> velocity=< 5,  2>
position=< 31990, -21060> velocity=<-3,  2>
position=< 10776, -52845> velocity=<-1,  5>
position=<-31636,  21324> velocity=< 3, -2>
position=<-10395, -42249> velocity=< 1,  4>
position=<-31591,  31912> velocity=< 3, -3>
position=<-10430, -21063> velocity=< 1,  2>
position=<-10422,  31912> velocity=< 1, -3>
position=< 10800, -42250> velocity=<-1,  4>
position=<-42199, -42248> velocity=< 4,  4>
position=< 31950,  42515> velocity=<-3, -4>
position=< 31982,  53104> velocity=<-3, -5>
position=< 31982, -52845> velocity=<-3,  5>
position=< 10752,  42510> velocity=<-1, -4>
position=< 31958, -21054> velocity=<-3,  2>
position=< 10752, -10467> velocity=<-1,  1>
position=< 31977, -31649> velocity=<-3,  3>
position=<-10438, -10460> velocity=< 1,  1>
position=<-20983,  10726> velocity=< 2, -1>
position=<-10385,  53109> velocity=< 1, -5>
position=< 42590, -52839> velocity=<-4,  5>
position=<-42180,  21317> velocity=< 4, -2>
position=<-31592, -21063> velocity=< 3,  2>
position=<-52782, -10466> velocity=< 5,  1>
position=<-31604, -42251> velocity=< 3,  4>
position=<-42198, -52848> velocity=< 4,  5>
position=<-31632,  42513> velocity=< 3, -4>
position=<-31601, -10468> velocity=< 3,  1>
position=<-31632,  21320> velocity=< 3, -2>
position=<-52773,  10725> velocity=< 5, -1>
position=<-42205, -10459> velocity=< 4,  1>
position=< 42533,  21322> velocity=<-4, -2>
position=< 53181,  53106> velocity=<-5, -5>
position=< 42533,  53102> velocity=<-4, -5>
position=< 53185, -42248> velocity=<-5,  4>
position=< 10797,  10730> velocity=<-1, -1>
position=< 53137,  21326> velocity=<-5, -2>
position=< 10776,  42508> velocity=<-1, -4>
position=< 21387, -42249> velocity=<-2,  4>
position=< 10805, -42250> velocity=<-1,  4>
position=<-20990,  42512> velocity=< 2, -4>
position=<-52801,  53102> velocity=< 5, -5>
position=< 42554,  53111> velocity=<-4, -5>
position=<-52809,  53106> velocity=< 5, -5>
position=<-31588, -42249> velocity=< 3,  4>
position=<-31595,  31921> velocity=< 3, -3>
position=< 42553, -10468> velocity=<-4,  1>
position=<-52818, -42249> velocity=< 5,  4>
position=<-20999, -42244> velocity=< 2,  4>
position=< 53127,  10731> velocity=<-5, -1>
position=< 21347, -52842> velocity=<-2,  5>
position=<-52794,  21323> velocity=< 5, -2>
position=<-42214,  31912> velocity=< 4, -3>
position=< 42558, -52847> velocity=<-4,  5>
position=<-10414, -42244> velocity=< 1,  4>
position=<-21025,  10728> velocity=< 2, -1>
position=<-10398,  10727> velocity=< 1, -1>
position=<-52794, -21056> velocity=< 5,  2>
position=< 10800,  53103> velocity=<-1, -5>
position=< 42556, -21059> velocity=<-4,  2>
position=<-31610,  42516> velocity=< 3, -4>
position=< 10780,  42507> velocity=<-1, -4>
position=< 31987, -52840> velocity=<-3,  5>
position=< 31952,  53102> velocity=<-3, -5>
position=<-21033, -42250> velocity=< 2,  4>
position=<-52773, -10466> velocity=< 5,  1>
position=<-42199,  42511> velocity=< 4, -4>
position=< 21355, -10467> velocity=<-2,  1>
position=<-21029, -10459> velocity=< 2,  1>
position=<-52775,  10727> velocity=< 5, -1>
position=<-21009,  31915> velocity=< 2, -3>
position=<-10389, -10464> velocity=< 1,  1>
position=< 53156,  31921> velocity=<-5, -3>
position=< 31995,  31920> velocity=<-3, -3>
position=<-31620, -52839> velocity=< 3,  5>
position=< 53158, -31654> velocity=<-5,  3>
position=<-42205,  42512> velocity=< 4, -4>
position=< 31950,  10725> velocity=<-3, -1>
position=<-10430, -42250> velocity=< 1,  4>
position=<-42206,  53102> velocity=< 4, -5>
position=< 42573,  10730> velocity=<-4, -1>
position=<-42178,  53111> velocity=< 4, -5>
position=<-10385,  53106> velocity=< 1, -5>
position=< 31955,  21320> velocity=<-3, -2>
position=< 42558,  10731> velocity=<-4, -1>
position=< 42585, -31656> velocity=<-4,  3>
position=< 42558, -21063> velocity=<-4,  2>
position=<-31612, -10461> velocity=< 3,  1>
position=<-10398,  31921> velocity=< 1, -3>
position=<-52765,  10724> velocity=< 5, -1>
position=< 42553,  53109> velocity=<-4, -5>
position=<-52810, -21061> velocity=< 5,  2>
position=<-42181, -31654> velocity=< 4,  3>
position=<-42175, -21058> velocity=< 4,  2>
position=<-10427, -10468> velocity=< 1,  1>
position=< 42529, -21055> velocity=<-4,  2>
position=< 10752,  53109> velocity=<-1, -5>
position=< 42561,  53107> velocity=<-4, -5>
position=< 31958, -31649> velocity=<-3,  3>
position=< 31968, -21063> velocity=<-3,  2>
position=<-10402, -52840> velocity=< 1,  5>
position=< 42565, -52848> velocity=<-4,  5>
position=< 53184,  31916> velocity=<-5, -3>
position=<-52797,  53111> velocity=< 5, -5>
position=< 42558, -52847> velocity=<-4,  5>
position=<-10394,  53106> velocity=< 1, -5>
position=<-42199, -42252> velocity=< 4,  4>
position=<-31592, -42250> velocity=< 3,  4>
position=< 10765,  53104> velocity=<-1, -5>
position=<-52794,  42509> velocity=< 5, -4>
position=< 21371,  53107> velocity=<-2, -5>
position=<-31616, -21063> velocity=< 3,  2>
position=<-42170,  31913> velocity=< 4, -3>
position=<-42199,  42513> velocity=< 4, -4>
position=< 10797,  53104> velocity=<-1, -5>
position=<-42175,  42511> velocity=< 4, -4>
position=<-20993,  21326> velocity=< 2, -2>
position=<-31588,  42513> velocity=< 3, -4>
position=<-20989,  31919> velocity=< 2, -3>
position=< 53140,  42509> velocity=<-5, -4>
position=< 53172, -10464> velocity=<-5,  1>
position=< 53128,  10724> velocity=<-5, -1>
position=<-31592,  42512> velocity=< 3, -4>
position=< 21347,  10724> velocity=<-2, -1>
position=<-52818, -21058> velocity=< 5,  2>
position=< 10794,  31912> velocity=<-1, -3>
position=<-10414, -52845> velocity=< 1,  5>
position=<-42219,  21326> velocity=< 4, -2>
position=< 31952,  31912> velocity=<-3, -3>
position=<-10413,  42507> velocity=< 1, -4>
position=< 21387, -31657> velocity=<-2,  3>
position=< 42557, -21054> velocity=<-4,  2>
position=<-20997,  53104> velocity=< 2, -5>
position=<-21014, -10468> velocity=< 2,  1>
position=<-10442,  10723> velocity=< 1, -1>
position=< 42572, -21054> velocity=<-4,  2>
position=<-20991, -42249> velocity=< 2,  4>
position=<-10435, -31649> velocity=< 1,  3>
position=< 10770, -10463> velocity=<-1,  1>
position=< 31958,  10730> velocity=<-3, -1>
position=<-21015, -42244> velocity=< 2,  4>
position=<-21039,  53111> velocity=< 2, -5>
position=<-31592,  31919> velocity=< 3, -3>
position=<-21022, -42249> velocity=< 2,  4>
position=< 31966,  10729> velocity=<-3, -1>
position=< 53167, -10468> velocity=<-5,  1>
position=<-42183,  21324> velocity=< 4, -2>
position=< 31955,  31915> velocity=<-3, -3>
position=< 42572,  53111> velocity=<-4, -5>
position=<-42170,  53102> velocity=< 4, -5>
position=< 42590, -10466> velocity=<-4,  1>
position=<-31609,  53106> velocity=< 3, -5>
position=< 42533, -10467> velocity=<-4,  1>
position=<-42215,  10731> velocity=< 4, -1>
position=<-42170,  10722> velocity=< 4, -1>
position=<-21004,  53102> velocity=< 2, -5>
position=< 53167,  10722> velocity=<-5, -1>
position=< 53132, -52844> velocity=<-5,  5>
position=< 42529, -21055> velocity=<-4,  2>
position=< 31960,  21322> velocity=<-3, -2>
position=<-42175,  42507> velocity=< 4, -4>
position=< 21342,  21326> velocity=<-2, -2>
position=<-52774,  31916> velocity=< 5, -3>
position=< 21355, -42248> velocity=<-2,  4>
position=< 53140,  10727> velocity=<-5, -1>
position=< 21368,  21317> velocity=<-2, -2>
position=< 42585, -42246> velocity=<-4,  4>
position=< 10805, -10459> velocity=<-1,  1>
position=< 21383, -52845> velocity=<-2,  5>
position=< 53133, -31649> velocity=<-5,  3>
position=<-31580, -31658> velocity=< 3,  3>
position=< 10792, -10463> velocity=<-1,  1>
position=< 10768, -52840> velocity=<-1,  5>
position=<-10420, -52848> velocity=< 1,  5>
position=<-42203,  31912> velocity=< 4, -3>
position=< 31974, -42245> velocity=<-3,  4>
position=<-52778,  42515> velocity=< 5, -4>
position=<-42170,  53106> velocity=< 4, -5>
position=<-10398, -10462> velocity=< 1,  1>
position=<-52799, -52848> velocity=< 5,  5>
position=<-42175,  42515> velocity=< 4, -4>
position=< 53142, -21059> velocity=<-5,  2>
position=< 10772,  42510> velocity=<-1, -4>
position=< 53137,  42516> velocity=<-5, -4>
position=<-42187, -21062> velocity=< 4,  2>
position=< 10805,  31914> velocity=<-1, -3>
position=< 21348,  10731> velocity=<-2, -1>
position=<-21023,  10726> velocity=< 2, -1>
position=< 31950, -31655> velocity=<-3,  3>
position=<-10386, -31654> velocity=< 1,  3>
position=<-21025, -21062> velocity=< 2,  2>
position=< 31966,  21318> velocity=<-3, -2>
position=<-52776,  31912> velocity=< 5, -3>
position=< 42569, -52841> velocity=<-4,  5>
position=<-31592,  10724> velocity=< 3, -1>
position=<-10429,  21317> velocity=< 1, -2>
position=<-31584,  10729> velocity=< 3, -1>
position=<-52818,  21320> velocity=< 5, -2>
position=<-21016, -10468> velocity=< 2,  1>
position=<-52791,  42507> velocity=< 5, -4>
position=<-42199,  42515> velocity=< 4, -4>
position=< 31982, -10467> velocity=<-3,  1>
position=<-10442, -52848> velocity=< 1,  5>
position=< 42565,  10726> velocity=<-4, -1>
position=<-42187, -10468> velocity=< 4,  1>
position=< 42561, -52841> velocity=<-4,  5>
position=< 42549,  31916> velocity=<-4, -3>
position=< 53180, -52842> velocity=<-5,  5>
position=< 21347,  21326> velocity=<-2, -2>
position=< 53151,  31921> velocity=<-5, -3>
position=< 31938,  10726> velocity=<-3, -1>
position=< 31942, -52848> velocity=<-3,  5>
position=< 10788, -10464> velocity=<-1,  1>
position=< 42590,  53103> velocity=<-4, -5>
position=< 53180, -21055> velocity=<-5,  2>
position=< 31935,  10731> velocity=<-3, -1>
position=<-52802, -31651> velocity=< 5,  3>
position=<-52776, -42249> velocity=< 5,  4>
position=< 53185, -21062> velocity=<-5,  2>
position=< 53159,  42511> velocity=<-5, -4>
position=< 10772,  10722> velocity=<-1, -1>
position=<-31593,  31912> velocity=< 3, -3>
position=< 10795,  21322> velocity=<-1, -2>
position=<-10436, -10459> velocity=< 1,  1>
position=< 21388, -42249> velocity=<-2,  4>
position=<-10405,  31921> velocity=< 1, -3>
position=<-52801,  42513> velocity=< 5, -4>
position=< 42573,  31915> velocity=<-4, -3>
position=< 21355,  53109> velocity=<-2, -5>
position=<-21025,  21323> velocity=< 2, -2>
position=< 53140, -10467> velocity=<-5,  1>
position=<-52767,  10726> velocity=< 5, -1>
position=< 21347,  31912> velocity=<-2, -3>
position=<-42183,  53107> velocity=< 4, -5>
position=< 31942, -10460> velocity=<-3,  1>
position=<-10401,  53102> velocity=< 1, -5>
position=< 21400, -10459> velocity=<-2,  1>
position=< 42579,  10722> velocity=<-4, -1>
position=< 31936,  42516> velocity=<-3, -4>
position=< 53128, -21056> velocity=<-5,  2>
position=<-20997, -42253> velocity=< 2,  4>
position=<-52777, -31658> velocity=< 5,  3>
position=<-10430,  21325> velocity=< 1, -2>
position=< 42545, -10464> velocity=<-4,  1>
position=<-21017, -52839> velocity=< 2,  5>
position=<-31584,  31912> velocity=< 3, -3>
position=< 53168, -21056> velocity=<-5,  2>
position=<-42199, -10459> velocity=< 4,  1>
position=< 10803,  53106> velocity=<-1, -5>
position=< 21371,  21317> velocity=<-2, -2>
position=<-21040, -42244> velocity=< 2,  4>
position=<-31577,  42511> velocity=< 3, -4>
position=< 42573,  21322> velocity=<-4, -2>
position=< 31990,  10728> velocity=<-3, -1>
position=< 53175, -52844> velocity=<-5,  5>
position=< 42561,  21323> velocity=<-4, -2>
position=<-42223, -52840> velocity=< 4,  5>
position=< 53152, -10468> velocity=<-5,  1>
position=<-52785,  53111> velocity=< 5, -5>
position=<-31632, -31656> velocity=< 3,  3>
position=<-21009, -10466> velocity=< 2,  1>
position=<-21025, -31651> velocity=< 2,  3>
position=< 10795,  53106> velocity=<-1, -5>
position=< 31978, -52844> velocity=<-3,  5>
position=<-52774, -21059> velocity=< 5,  2>
position=<-42211, -42253> velocity=< 4,  4>
position=<-31580,  53107> velocity=< 3, -5>
position=< 53172,  53102> velocity=<-5, -5>
position=<-52773,  42508> velocity=< 5, -4>
position=<-10398, -21054> velocity=< 1,  2>
position=<-31607, -10466> velocity=< 3,  1>
position=<-52801,  53108> velocity=< 5, -5>
position=<-52813, -21054> velocity=< 5,  2>
position=<-31600,  42511> velocity=< 3, -4>
position=<-31631,  21317> velocity=< 3, -2>
position=< 42532, -31649> velocity=<-4,  3>`
