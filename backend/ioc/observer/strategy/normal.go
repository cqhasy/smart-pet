package strategy

type NormalStrategy struct {
}

func (n *NormalStrategy) Calc(subX, subY int) (objX, objY int) {
	return subX, subY
}
