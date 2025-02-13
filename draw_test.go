package mpb

import (
	"bytes"
	"testing"
	"unicode/utf8"
)

func TestDraw(t *testing.T) {
	// key is termWidth
	testSuite := map[int][]struct {
		style    BarStyleComposer
		name     string
		total    int64
		current  int64
		refill   int64
		barWidth int
		trim     bool
		want     string
	}{
		0: {
			{
				name:    "t,c{60,20}",
				total:   60,
				current: 20,
				want:    "",
			},
			{
				name:    "t,c{60,20}trim",
				total:   60,
				current: 20,
				trim:    true,
				want:    "",
			},
		},
		1: {
			{
				name:    "t,c{60,20}",
				total:   60,
				current: 20,
				want:    "",
			},
			{
				name:    "t,c{60,20}trim",
				total:   60,
				current: 20,
				trim:    true,
				want:    "",
			},
		},
		2: {
			{
				name:    "t,c{60,20}",
				total:   60,
				current: 20,
				want:    "  ",
			},
			{
				name:    "t,c{60,20}trim",
				total:   60,
				current: 20,
				trim:    true,
				want:    "[]",
			},
		},
		3: {
			{
				name:    "t,c{60,20}",
				total:   60,
				current: 20,
				want:    "  ",
			},
			{
				name:    "t,c{60,20}trim",
				total:   60,
				current: 20,
				trim:    true,
				want:    "[-]",
			},
		},
		4: {
			{
				name:    "t,c{60,20}",
				total:   60,
				current: 20,
				want:    " [] ",
			},
			{
				name:    "t,c{60,20}trim",
				total:   60,
				current: 20,
				trim:    true,
				want:    "[>-]",
			},
		},
		5: {
			{
				name:    "t,c{60,20}",
				total:   60,
				current: 20,
				want:    " [-] ",
			},
			{
				name:    "t,c{60,20}trim",
				total:   60,
				current: 20,
				trim:    true,
				want:    "[>--]",
			},
		},
		6: {
			{
				name:    "t,c{60,20}",
				total:   60,
				current: 20,
				want:    " [>-] ",
			},
			{
				name:    "t,c{60,20}trim",
				total:   60,
				current: 20,
				trim:    true,
				want:    "[>---]",
			},
		},
		7: {
			{
				name:    "t,c{60,20}",
				total:   60,
				current: 20,
				want:    " [>--] ",
			},
			{
				name:    "t,c{60,20}trim",
				total:   60,
				current: 20,
				trim:    true,
				want:    "[=>---]",
			},
		},
		8: {
			{
				name:    "t,c{60,20}",
				total:   60,
				current: 20,
				want:    " [>---] ",
			},
			{
				name:    "t,c{60,20}trim",
				total:   60,
				current: 20,
				trim:    true,
				want:    "[=>----]",
			},
		},
		80: {
			{
				name:    "t,c{60,20}",
				total:   60,
				current: 20,
				want:    " [========================>---------------------------------------------------] ",
			},
			{
				name:    "t,c{60,20}trim",
				total:   60,
				current: 20,
				trim:    true,
				want:    "[=========================>----------------------------------------------------]",
			},
			{
				name:     "t,c,bw{60,20,60}",
				total:    60,
				current:  20,
				barWidth: 60,
				want:     " [==================>---------------------------------------] ",
			},
			{
				name:     "t,c,bw{60,20,60}trim",
				total:    60,
				current:  20,
				barWidth: 60,
				trim:     true,
				want:     "[==================>---------------------------------------]",
			},
		},
		99: {
			{
				style:   BarStyle().Lbound("[").Filler("の").Tip("だ").Padding("つ").Rbound("]"),
				name:    "[のだつ] t,c{99,1}",
				total:   99,
				current: 1,
				want:    " [だつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつ…] ",
			},
			{
				style:   BarStyle().Lbound("[").Filler("の").Tip("だ").Padding("つ").Rbound("]"),
				name:    "[のだつ] t,c{99,2}",
				total:   99,
				current: 2,
				want:    " [だつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつ…] ",
			},
			{
				style:   BarStyle().Lbound("[").Filler("の").Tip("だ").Padding("つ").Rbound("]"),
				name:    "[のだつ] t,c{99,3}",
				total:   99,
				current: 3,
				want:    " [だつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつ…] ",
			},
			{
				style:   BarStyle().Lbound("[").Filler("の").Tip("だ").Padding("つ").Rbound("]"),
				name:    "[のだつ] t,c{99,4}",
				total:   99,
				current: 4,
				want:    " [のだつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつ…] ",
			},
			{
				style:   BarStyle().Lbound("[").Filler("の").Tip("だ").Padding("つ").Rbound("]"),
				name:    "[のだつ] t,c{99,5}",
				total:   99,
				current: 5,
				want:    " [のだつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつ…] ",
			},
			{
				style:   BarStyle().Lbound("[").Filler("の").Tip("だ").Padding("つ").Rbound("]"),
				name:    "[のだつ] t,c{99,6}",
				total:   99,
				current: 6,
				want:    " [ののだつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつ…] ",
			},
			{
				style:   BarStyle().Lbound("[").Filler("の").Tip("だ").Padding("つ").Rbound("]").Reverse(),
				name:    "[のだつ] t,c{99,6}rev",
				total:   99,
				current: 6,
				want:    " […つつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつつだのの] ",
			},
		},
		100: {
			{
				name:    "t,c{100,0}",
				total:   100,
				current: 0,
				want:    " [------------------------------------------------------------------------------------------------] ",
			},
			{
				name:    "t,c{100,0}trim",
				total:   100,
				current: 0,
				trim:    true,
				want:    "[--------------------------------------------------------------------------------------------------]",
			},
			{
				name:    "t,c{100,1}",
				total:   100,
				current: 1,
				want:    " [>-----------------------------------------------------------------------------------------------] ",
			},
			{
				name:    "t,c{100,1}trim",
				total:   100,
				current: 1,
				trim:    true,
				want:    "[>-------------------------------------------------------------------------------------------------]",
			},
			{
				name:    "t,c{100,99}",
				total:   100,
				current: 99,
				want:    " [==============================================================================================>-] ",
			},
			{
				name:    "t,c{100,99}trim",
				total:   100,
				current: 99,
				trim:    true,
				want:    "[================================================================================================>-]",
			},
			{
				name:    "t,c{100,100}",
				total:   100,
				current: 100,
				want:    " [================================================================================================] ",
			},
			{
				name:    "t,c{100,100}trim",
				total:   100,
				current: 100,
				trim:    true,
				want:    "[==================================================================================================]",
			},
			{
				name:    "t,c,r{100,100,100}trim",
				total:   100,
				current: 100,
				refill:  100,
				trim:    true,
				want:    "[++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++]",
			},
			{
				name:    "t,c{100,33}",
				total:   100,
				current: 33,
				want:    " [===============================>----------------------------------------------------------------] ",
			},
			{
				name:    "t,c{100,33}trim",
				total:   100,
				current: 33,
				trim:    true,
				want:    "[===============================>------------------------------------------------------------------]",
			},
			{
				style:   BarStyle().Tip("<").Reverse(),
				name:    "t,c{100,33}trim,rev",
				total:   100,
				current: 33,
				trim:    true,
				want:    "[------------------------------------------------------------------<===============================]",
			},
			{
				name:    "t,c,r{100,33,33}",
				total:   100,
				current: 33,
				refill:  33,
				want:    " [+++++++++++++++++++++++++++++++>----------------------------------------------------------------] ",
			},
			{
				name:    "t,c,r{100,33,33}trim",
				total:   100,
				current: 33,
				refill:  33,
				trim:    true,
				want:    "[+++++++++++++++++++++++++++++++>------------------------------------------------------------------]",
			},
			{
				style:   BarStyle().Tip("<").Reverse(),
				name:    "t,c,r{100,33,33}trim,rev",
				total:   100,
				current: 33,
				refill:  33,
				trim:    true,
				want:    "[------------------------------------------------------------------<+++++++++++++++++++++++++++++++]",
			},
			{
				name:    "t,c,r{100,40,33}",
				total:   100,
				current: 40,
				refill:  33,
				want:    " [++++++++++++++++++++++++++++++++=====>----------------------------------------------------------] ",
			},
			{
				name:    "t,c,r{100,40,33}trim",
				total:   100,
				current: 40,
				refill:  33,
				trim:    true,
				want:    "[++++++++++++++++++++++++++++++++======>-----------------------------------------------------------]",
			},
			{
				style:   BarStyle().Tip("<").Reverse(),
				name:    "t,c,r{100,40,33},rev",
				total:   100,
				current: 40,
				refill:  33,
				want:    " [----------------------------------------------------------<=====++++++++++++++++++++++++++++++++] ",
			},
			{
				style:   BarStyle().Tip("<").Reverse(),
				name:    "t,c,r{100,40,33}trim,rev",
				total:   100,
				current: 40,
				refill:  33,
				trim:    true,
				want:    "[-----------------------------------------------------------<======++++++++++++++++++++++++++++++++]",
			},
		},
		197: {
			{
				name:     "t,c,r{97486999,2805950,2805483}trim",
				total:    97486999,
				current:  2805950,
				refill:   2805483,
				barWidth: 60,
				trim:     true,
				want:     "[+>--------------------------------------------------------]",
			},
		},
	}

	var tmpBuf bytes.Buffer
	for tw, cases := range testSuite {
		for _, tc := range cases {
			if tc.style == nil {
				tc.style = BarStyle()
			}
			s := newTestState(NewBarFiller(tc.style))
			s.reqWidth = tc.barWidth
			s.total = tc.total
			s.current = tc.current
			s.trimSpace = tc.trim
			s.refill = tc.refill
			tmpBuf.Reset()
			tmpBuf.ReadFrom(s.draw(newStatistics(tw, s)))
			by := tmpBuf.Bytes()

			got := string(by[:len(by)-1])
			if !utf8.ValidString(got) {
				t.Fail()
			}
			if got != tc.want {
				t.Errorf("termWidth:%d %q want: %q %d, got: %q %d\n", tw, tc.name, tc.want, utf8.RuneCountInString(tc.want), got, utf8.RuneCountInString(got))
			}
		}
	}
}

func newTestState(filler BarFiller) *bState {
	s := &bState{
		filler: filler,
		bufP:   new(bytes.Buffer),
		bufB:   new(bytes.Buffer),
		bufA:   new(bytes.Buffer),
	}
	return s
}
