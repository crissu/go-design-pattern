package register

type RegisterInterface interface {
	run() string
	smile() string
	sleep() string
}

var OneRegister RegisterInterface

func NewRegister(cla string) RegisterInterface {
	if cla == "cat" {
		cat := new(cat)
		cat.name = "haloketti"
		cat.age = 1
		cat.color = "blue"
		cat.description = "nice"
		return cat
	}
	return nil
}

func Register(register RegisterInterface) {
	OneRegister = register
}

func Run(){
	OneRegister.run()
}

func Smile() {
	OneRegister.smile()
}

func Sleep() {
	OneRegister.sleep()
}
