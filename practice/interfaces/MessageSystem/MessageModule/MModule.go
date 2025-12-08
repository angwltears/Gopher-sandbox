package mmodule

type Msg interface {
	Send()
}

func SendAll(m []Msg) {
	for _, v := range m {
		v.Send()
	}
}
