package main

import spx "github.com/goplus/spx"

type index struct {
	spx.Game
	Monkey Monkey
}
type Monkey struct {
	spx.Sprite
	*index
}

func (this *index) MainEntry() {
//line D:\workplace\goplus\QinSpx\index.gmx:4
	spx.Gopt_Game_Run(this, "res", &spx.Config{Title: "Qin Spx(By Go+ Engine)"})
}
func main() {
	spx.Gopt_Game_Main(new(index))
}

const oneStep = 20

func (this *Monkey) turnOrStep(dir float64) {
//line D:\workplace\goplus\QinSpx\Monkey.spx:4
	if this.Heading() == dir {
//line D:\workplace\goplus\QinSpx\Monkey.spx:6
		this.Step__0(oneStep)
	} else {
//line D:\workplace\goplus\QinSpx\Monkey.spx:8
		this.TurnTo(dir)
	}
}
func (this *Monkey) Main() {
//line D:\workplace\goplus\QinSpx\Monkey.spx:12
	this.OnCloned__1(func() {
//line D:\workplace\goplus\QinSpx\Monkey.spx:13
		this.SetXYpos(this.MouseX(), this.MouseY())
//line D:\workplace\goplus\QinSpx\Monkey.spx:14
		this.Show()
//line D:\workplace\goplus\QinSpx\Monkey.spx:15
		for {
			spx.Sched()
//line D:\workplace\goplus\QinSpx\Monkey.spx:16
			this.Wait(0.04)
//line D:\workplace\goplus\QinSpx\Monkey.spx:17
			this.Step__0(10)
//line D:\workplace\goplus\QinSpx\Monkey.spx:18
			if this.Touching(spx.Edge) {
//line D:\workplace\goplus\QinSpx\Monkey.spx:19
				this.Destroy()
			}
		}
	})
//line D:\workplace\goplus\QinSpx\Monkey.spx:24
	this.OnStart(func() {
//line D:\workplace\goplus\QinSpx\Monkey.spx:25
		this.Say("Hello Qin")
//line D:\workplace\goplus\QinSpx\Monkey.spx:29
		for {
			spx.Sched()
//line D:\workplace\goplus\QinSpx\Monkey.spx:30
			this.Wait(1)
//line D:\workplace\goplus\QinSpx\Monkey.spx:31
			spx.Gopt_Sprite_Clone__0(this)
		}
	})
//line D:\workplace\goplus\QinSpx\Monkey.spx:35
	this.OnKey__0(spx.KeyUp, func() {
//line D:\workplace\goplus\QinSpx\Monkey.spx:37
		this.turnOrStep(0)
	})
//line D:\workplace\goplus\QinSpx\Monkey.spx:40
	this.OnKey__0(spx.KeyRight, func() {
//line D:\workplace\goplus\QinSpx\Monkey.spx:42
		this.turnOrStep(90)
	})
//line D:\workplace\goplus\QinSpx\Monkey.spx:45
	this.OnKey__0(spx.KeyDown, func() {
//line D:\workplace\goplus\QinSpx\Monkey.spx:47
		this.turnOrStep(180)
	})
//line D:\workplace\goplus\QinSpx\Monkey.spx:50
	this.OnKey__0(spx.KeyLeft, func() {
//line D:\workplace\goplus\QinSpx\Monkey.spx:52
		this.turnOrStep(-90)
	})
//line D:\workplace\goplus\QinSpx\Monkey.spx:55
	this.OnKey__0(spx.KeyA, func() {
//line D:\workplace\goplus\QinSpx\Monkey.spx:56
		this.Turn(spx.Left)
	})
//line D:\workplace\goplus\QinSpx\Monkey.spx:59
	this.OnKey__0(spx.KeyD, func() {
//line D:\workplace\goplus\QinSpx\Monkey.spx:60
		this.Turn(spx.Right)
	})
//line D:\workplace\goplus\QinSpx\Monkey.spx:63
	this.OnKey__0(spx.KeyW, func() {
//line D:\workplace\goplus\QinSpx\Monkey.spx:64
		this.Turn(-45)
	})
//line D:\workplace\goplus\QinSpx\Monkey.spx:67
	this.OnKey__0(spx.KeyS, func() {
//line D:\workplace\goplus\QinSpx\Monkey.spx:68
		this.Turn(45)
	})
//line D:\workplace\goplus\QinSpx\Monkey.spx:73
	this.OnClick(func() {
//line D:\workplace\goplus\QinSpx\Monkey.spx:74
		this.Say("press the arrorw keys", 1)
	})
}
