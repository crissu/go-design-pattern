package register

func test() string {
	cat := NewRegister("cat")
	Register(cat)
	return OneRegister.smile() + " - " + OneRegister.sleep() + " - " + OneRegister.run()
}
