package js

type ElseIfChain struct {
	q Ctx
}

func (e ElseIfChain) Else(do Script) {
	q := e.q
	q(" else {")
	do(q)
	q("}")
}

func (q Ctx) If(condition AnyBool, do Script) ElseIfChain {
	q("if(")
	q(condition.GetBool())
	q(") {")
	q(do)
	q("}")

	return ElseIfChain{q}
}

func If(condition AnyBool, do Script) Script {
	return func(q Ctx) {
		q.If(condition, do)
	}
}

func (s Script) Else(do Script) Script {
	return func(q Ctx) {
		s(q)
		q("else {")
		do(q)
		q("}")
	}
}