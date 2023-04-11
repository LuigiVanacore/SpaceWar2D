package input

//import "github.com/LuigiVanacore/AirWars2D/historia_engine/data_structure"
//
//type MappedInput struct {
//	actions data_structure.Set[Action]
//	states  data_structure.Set[State]
//	ranges  map[Range]float64
//}
//
//func NewMappedInput() *MappedInput {
//	return &MappedInput{
//		actions: data_structure.NewSet[Action](),
//		states:  data_structure.NewSet[State](),
//		ranges:  make(map[Range]float64)}
//}
//
//func (m *MappedInput) EatAction(action Action) {
//	m.actions.Remove(action)
//}
//
//func (m *MappedInput) EatState(state State) {
//	m.states.Remove(state)
//}
//
//func (m *MappedInput) EatRange(axisRange Range) {
//	delete(m.ranges, axisRange)
//}
//
//func (m *MappedInput) Clear() {
//	m.actions.Clear()
//	m.states.Clear()
//	m.ranges = nil
//}
//
//func (m *MappedInput) InsertAction(action Action) {
//	m.actions.Add(action)
//}
//
//func (m *MappedInput) InsertState(state State) {
//	m.states.Add(state)
//}
