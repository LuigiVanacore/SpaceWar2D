package input

//
//type InputContext struct {
//	actionMap map[RawInputButton]Action
//	stateMap  map[RawInputButton]State
//	rangeMap  map[RawInputAxis]Range
//
//	sensitivityMap map[Range]float64
//	conversions    *RangeConverter
//}
//
//func (i InputContext) MapButtonToAction(button RawInputButton) (bool, Action) {
//	action, ok := i.actionMap[button]
//	return ok, action
//}
//
//func (i InputContext) MapButtonToState(button RawInputButton) (bool, State) {
//	state, ok := i.stateMap[button]
//	return ok, state
//}
//
//func (i InputContext) MapAxisToRange(axis RawInputAxis) (bool, Range) {
//	axisRange, ok := i.rangeMap[axis]
//	return ok, axisRange
//}
//
//func (i InputContext) GetSensitivity(_range Range) float64 {
//	sensitivity, ok := i.sensitivityMap[_range]
//	if ok == false {
//		return 1.0
//	}
//	return sensitivity
//}
//
//func (i InputContext) GetConversions() *RangeConverter {
//	return i.conversions
//}
