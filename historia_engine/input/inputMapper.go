package input

//
//type InputMapper struct {
//	inputContext  map[string]*InputContext
//	activeContext []*InputContext
//
//	callbackMap map[int]InputCallback
//
//	currentMappedInput MappedInput
//}
//
//type InputCallback func(input *MappedInput)
//
//func (i *InputMapper) Clear() {
//	i.currentMappedInput.Clear()
//}
//
//func (i *InputMapper) SetRawButtonState(button RawInputButton, pressed bool, previouslypressed bool) {
//
//	if pressed && !previouslypressed {
//		ok, action := i.mapButtonToAction(button)
//		if ok {
//			i.currentMappedInput.InserAction(button, action)
//			return
//		}
//	}
//
//	if pressed {
//		ok, state := i.mapButtonToState(button)
//		if ok {
//			i.currentMappedInput.InsertState(button, state)
//			return
//		}
//	}
//
//	i.mapAndEatButton(button)
//
//}
//
//func (i *InputMapper) SetRawAxisValue(axis RawInputAxis, value float64) {
//
//	for _, context := range i.activeContext {
//		ok, _range := context.MapAxisToRange(axis)
//		if ok {
//			i.currentMappedInput.rangeMap[axis] = context.GetConversions().Convert(_range, value*context.GetSensitivity(_range))
//			break
//		}
//	}
//}
//
//func (i *InputMapper) Dispatch() {
//
//	for _, value := range i.callbackMap {
//		value(&i.currentMappedInput)
//	}
//}
//
//func (i *InputMapper) AddCallback(callback InputCallback, priority int) {
//	i.callbackMap[priority] = callback
//}
//
//func (i *InputMapper) PushContext(name string) {
//	context, _ := i.inputContext[name]
//	i.activeContext = append(i.activeContext, context)
//}
//
//func (i *InputMapper) PopContext() {
//	if len(i.activeContext) != 0 {
//		index := 0
//		i.activeContext = append(i.activeContext[:index], i.activeContext[index+1:]...)
//	}
//}
//
//func (i *InputMapper) mapButtonToAction(button RawInputButton) (bool, Action) {
//
//	for _, context := range i.activeContext {
//		ok, action := context.MapButtonToAction(button)
//		if ok {
//			return true, action
//		}
//	}
//	return false, 0
//}
//
//func (i *InputMapper) mapButtonToState(button RawInputButton) (bool, State) {
//
//	for _, context := range i.activeContext {
//		ok, state := context.MapButtonToState(button)
//		if ok {
//			return true, state
//		}
//	}
//	return false, 0
//}
//
//func (i *InputMapper) mapAndEatButton(button RawInputButton) {
//
//	ok, action := i.mapButtonToAction(button)
//	if ok {
//		i.currentMappedInput.EatAction(action)
//	}
//
//	ok, state := i.mapButtonToState(button)
//	if ok {
//		i.currentMappedInput.EatState(state)
//	}
//}
