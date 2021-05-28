package main

import "fmt"

/////////////////// BATTERY CLASS ///////////////////////
type Battery struct {
	ID                        int
	status                    string
	columnsList               []Column
	floorRequestButtonsList   []FloorRequestButtons
	amountOfColumns           int
	amountOfFloors            int
	amountOfElevatorPerColumn int
	amountOfBasements         int
}

func NewBattery(_id int, _amountOfColumns int, _amountOfFloors int, _amountOfBasements int, _amountOfElevatorPerColumn int) Battery {
	return Battery{
		ID:                        _id,
		status:                    "idle",
		columnsList:               []Column{},
		floorRequestButtonsList:   []FloorRequestButtons{},
		amountOfColumns:           _amountOfColumns,
		amountOfFloors:            _amountOfFloors,
		amountOfElevatorPerColumn: _amountOfElevatorPerColumn,
		amountOfBasements:         _amountOfBasements,
	}
}

/////////////////// CREATE COLUMN ///////////////////////

func (m Battery) createColumn(_amountOfColumns int, _amountOfFloors int, _amountOfElevators int) []Column {

	for i := 0; i < _amountOfColumns; i++ {
		column := Column{ID: i, amountOfFloors: _amountOfFloors, amountOfElevators: _amountOfElevators}
		m.columnsList = append(m.columnsList, column)
		m.columnsList[0].isBasement = true
		m.columnsList[i].elevatorsList = append(m.columnsList[i].elevatorsList, m.columnsList[i].createElevator(_amountOfFloors, m.amountOfElevatorPerColumn)...)
		m.columnsList[i].createElevator(_amountOfFloors, _amountOfElevators)
		if m.columnsList[i].isBasement == false {
			var servedFloorMin = (i - 1) * (_amountOfFloors / (m.amountOfColumns - 1))
			var servedFloorMax = i * (_amountOfFloors / (m.amountOfColumns - 1))
			m.columnsList[i].servedFloors = append(m.columnsList[i].servedFloors, 0)
			m.columnsList[i].servedFloors = append(m.columnsList[i].servedFloors, servedFloorMin)
			m.columnsList[i].servedFloors = append(m.columnsList[i].servedFloors, servedFloorMax)
			/* fmt.Printf("%+v", m.columnsList[i]) */
		} else {
			var servedFloorMin = i
			var servedFloorMax = i - m.amountOfBasements
			m.columnsList[i].servedFloors = append(m.columnsList[i].servedFloors, 0)
			m.columnsList[i].servedFloors = append(m.columnsList[i].servedFloors, -servedFloorMin)
			m.columnsList[i].servedFloors = append(m.columnsList[i].servedFloors, servedFloorMax)
		}
	}
	return m.columnsList
}

/////////////////// find best column ///////////////////////

func (j Battery) findBestColumn(_requestedFloor int, columnsList []Column) int {
	for i := 0; i < len(columnsList); i++ {
		if columnsList[i].servedFloors[1] <= _requestedFloor && columnsList[i].servedFloors[2] >= _requestedFloor || columnsList[i].servedFloors[1] >= _requestedFloor && columnsList[i].servedFloors[2] <= _requestedFloor {
			var selectedColumn = i
			return selectedColumn
		}
	}
	return 0
}

/////////////////// assignElevator ///////////////////////

func (r Battery) assignElevator(_requestedFloor int, _direction string, _requestedFloorAfter int, _directionAfter string) {
	// create column & elevator
	var createdColumn = r.createColumn(r.amountOfColumns, r.amountOfFloors, r.amountOfElevatorPerColumn)
	// find best column
	var selectedColumnNumber = r.findBestColumn(_requestedFloor, createdColumn)
	var selectedColumn = createdColumn[selectedColumnNumber]
	fmt.Printf("\n%s %d %s", "Column : ", selectedColumnNumber, " is selected")

	// Set floor on scenario
	selectedColumn.elevatorsList[0].currentFloor = 20
	selectedColumn.elevatorsList[1].currentFloor = 3
	selectedColumn.elevatorsList[2].currentFloor = 13
	selectedColumn.elevatorsList[3].currentFloor = 15
	selectedColumn.elevatorsList[4].currentFloor = 6

	// find best elevator
	selectedColumn.status = "online"

	for i := range selectedColumn.elevatorsList {
		if selectedColumn.elevatorsList[i].currentFloor == _requestedFloor {
			selectedColumn.elevatorsList[i].floorRequestList = append(selectedColumn.elevatorsList[i].floorRequestList, _requestedFloor)
			selectedColumn.status = "busy"
			break
		}
	}
	if selectedColumn.status == "online" {
		for i := range selectedColumn.elevatorsList {
			if selectedColumn.elevatorsList[i].currentFloor-_requestedFloor >= -1 && selectedColumn.elevatorsList[i].currentFloor-_requestedFloor <= 1 {
				selectedColumn.elevatorsList[i].floorRequestList = append(selectedColumn.elevatorsList[i].floorRequestList, _requestedFloor)
				selectedColumn.status = "busy"
				break
			}
		}
	}
	if selectedColumn.status == "online" {
		for i := range selectedColumn.elevatorsList {
			if len(selectedColumn.elevatorsList[i].floorRequestList) == 0 {
				if selectedColumn.elevatorsList[i].currentFloor-_requestedFloor > -2 && selectedColumn.elevatorsList[i].currentFloor-_requestedFloor < 2 {
					selectedColumn.elevatorsList[i].floorRequestList = append(selectedColumn.elevatorsList[i].floorRequestList, _requestedFloor)
					selectedColumn.status = "busy"
					break
				}
			}
		}
	}
	if selectedColumn.status == "online" {
		for i := range selectedColumn.elevatorsList {
			if len(selectedColumn.elevatorsList[i].floorRequestList) == 0 {
				if selectedColumn.elevatorsList[i].currentFloor-_requestedFloor > -3 && selectedColumn.elevatorsList[i].currentFloor-_requestedFloor < 3 {
					selectedColumn.elevatorsList[i].floorRequestList = append(selectedColumn.elevatorsList[i].floorRequestList, _requestedFloor)
					selectedColumn.status = "busy"
					break
				}
			}
		}
	}
	if selectedColumn.status == "online" {
		for i := range selectedColumn.elevatorsList {
			if len(selectedColumn.elevatorsList[i].floorRequestList) == 0 {
				if selectedColumn.elevatorsList[i].currentFloor-_requestedFloor > -4 && selectedColumn.elevatorsList[i].currentFloor-_requestedFloor < 4 {
					selectedColumn.elevatorsList[i].floorRequestList = append(selectedColumn.elevatorsList[i].floorRequestList, _requestedFloor)
					selectedColumn.status = "busy"
					break
				}
			}
		}
	}
	if selectedColumn.status == "online" {
		for i := range selectedColumn.elevatorsList {
			if len(selectedColumn.elevatorsList[i].floorRequestList) == 0 {
				if selectedColumn.elevatorsList[i].currentFloor-_requestedFloor > -5 && selectedColumn.elevatorsList[i].currentFloor-_requestedFloor < 5 {
					selectedColumn.elevatorsList[i].floorRequestList = append(selectedColumn.elevatorsList[i].floorRequestList, _requestedFloor)
					selectedColumn.status = "busy"
					break
				}
			}
		}
	}
	if selectedColumn.status == "online" {
		for i := range selectedColumn.elevatorsList {
			if len(selectedColumn.elevatorsList[i].floorRequestList) == 0 {
				if selectedColumn.elevatorsList[i].currentFloor-_requestedFloor > -10 && selectedColumn.elevatorsList[i].currentFloor-_requestedFloor < 10 {
					selectedColumn.elevatorsList[i].floorRequestList = append(selectedColumn.elevatorsList[i].floorRequestList, _requestedFloor)
					selectedColumn.status = "busy"
					break
				}
			}
		}
	}
	if selectedColumn.status == "online" {
		for i := range selectedColumn.elevatorsList {
			if len(selectedColumn.elevatorsList[i].floorRequestList) == 0 {
				if selectedColumn.elevatorsList[i].currentFloor-_requestedFloor > -20 && selectedColumn.elevatorsList[i].currentFloor-_requestedFloor < 20 {
					selectedColumn.elevatorsList[i].floorRequestList = append(selectedColumn.elevatorsList[i].floorRequestList, _requestedFloor)
					selectedColumn.status = "busy"
					break
				}
			}
		}
	}

	/// mouvement

	for i := 0; i < r.amountOfElevatorPerColumn; i++ {
		if len(selectedColumn.elevatorsList[i].floorRequestList) > 0 {
			selectedColumn.elevatorsList[i].Door.status = "closed"
			selectedColumn.elevatorsList[i].status = "moving"
			if selectedColumn.elevatorsList[i].currentFloor < _requestedFloor {
				selectedColumn.elevatorsList[i].direction = "up"
				fmt.Printf("\n%s %d %s %s %s", "Elevator ", selectedColumn.elevatorsList[i].ID, " is ", selectedColumn.elevatorsList[i].status, selectedColumn.elevatorsList[i].direction)
				for x := selectedColumn.elevatorsList[i].currentFloor; x < _requestedFloor; x++ {
					selectedColumn.elevatorsList[i].currentFloor = selectedColumn.elevatorsList[i].currentFloor + 1
					fmt.Printf("\n%s %d", "Floor : ", selectedColumn.elevatorsList[i].currentFloor)
				}
				//// ouverture des portes
				selectedColumn.elevatorsList[i].Door.status = "opened"
				fmt.Printf("\nThe door is : " + selectedColumn.elevatorsList[i].Door.status)
			} else if selectedColumn.elevatorsList[i].currentFloor > _requestedFloor {
				selectedColumn.elevatorsList[i].direction = "down"
				fmt.Printf("\n%s %d %s %s %s", "Elevator ", selectedColumn.elevatorsList[i].ID, " is ", selectedColumn.elevatorsList[i].status, selectedColumn.elevatorsList[i].direction)
				for x := selectedColumn.elevatorsList[i].currentFloor; x > _requestedFloor; x-- {
					selectedColumn.elevatorsList[i].currentFloor = selectedColumn.elevatorsList[i].currentFloor - 1
					fmt.Printf("\n%s %d", "Floor : ", selectedColumn.elevatorsList[i].currentFloor)
				}
				//// ouverture des portes
				selectedColumn.elevatorsList[i].Door.status = "opened"
				fmt.Printf("\nThe door is : " + selectedColumn.elevatorsList[i].Door.status)
			} else {
				//// ouverture des portes
				selectedColumn.elevatorsList[i].Door.status = "opened"
				fmt.Printf("\nThe door is : " + selectedColumn.elevatorsList[i].Door.status)
			}
		}
	}

	selectedColumn.requestElevator(_requestedFloorAfter, _directionAfter)
	selectedColumn.status = "online"
	/* fmt.Printf("%+v", selectedColumn.elevatorsList) */

	/* fmt.Printf("%s %d", "ici", _requestedFloor)  exemple pour associer string + int */

}

/////////////////// COLUMN CLASS ///////////////////////

type Column struct {
	ID                int
	status            string
	servedFloors      []int
	isBasement        bool
	elevatorsList     []Elevator
	callButtonsList   []int
	amountOfElevators int
	amountOfFloors    int
}

func NewColumn(_id int, _amountOfFloors int, _amountOfElevators int) Column {
	return Column{
		ID:                _id,
		status:            "online",
		servedFloors:      []int{},
		isBasement:        false,
		elevatorsList:     []Elevator{},
		callButtonsList:   []int{},
		amountOfElevators: _amountOfElevators,
		amountOfFloors:    _amountOfFloors,
	}
}

/////////////////// CREATE ELEVATOR ///////////////////////

func (r Column) createElevator(_amountOfFloors int, _amountOfElevators int) []Elevator {
	for i := 0; i < _amountOfElevators; i++ {
		var elevator = Elevator{ID: i, amountOfFloors: _amountOfFloors}
		r.elevatorsList = append(r.elevatorsList, elevator)

	}
	return r.elevatorsList

}

/////////////////// requestElevator ///////////////////////

func (s Column) requestElevator(_requestedFloor int, _direction string) {
	for i := 0; i < s.amountOfElevators; i++ {
		if len(s.elevatorsList[i].floorRequestList) > 0 {
			s.elevatorsList[i].Door.status = "closed"
			s.elevatorsList[i].status = "moving"
			fmt.Printf("\nThe door is : " + s.elevatorsList[i].Door.status)
			if s.elevatorsList[i].currentFloor < _requestedFloor {
				s.elevatorsList[i].direction = "up"
				fmt.Printf("\n%s %d %s %s %s", "Elevator ", s.elevatorsList[i].ID, " is ", s.elevatorsList[i].status, s.elevatorsList[i].direction)
				for x := s.elevatorsList[i].currentFloor; x < _requestedFloor; x++ {
					s.elevatorsList[i].currentFloor = s.elevatorsList[i].currentFloor + 1
					fmt.Printf("\n%s %d", "Floor : ", s.elevatorsList[i].currentFloor)
				}
				//// ouverture des portes
				s.elevatorsList[i].Door.status = "opened"
				fmt.Printf("\nThe door is : " + s.elevatorsList[i].Door.status)
			} else {
				s.elevatorsList[i].direction = "down"
				fmt.Printf("\n%s %d %s %s %s", "Elevator ", s.elevatorsList[i].ID, " is ", s.elevatorsList[i].status, s.elevatorsList[i].direction)
				for x := s.elevatorsList[i].currentFloor; x > _requestedFloor; x-- {
					s.elevatorsList[i].currentFloor = s.elevatorsList[i].currentFloor - 1
					fmt.Printf("\n%s %d", "Floor : ", s.elevatorsList[i].currentFloor)
				}
				//// ouverture des portes
				s.elevatorsList[i].Door.status = "opened"
				fmt.Printf("\nThe door is : " + s.elevatorsList[i].Door.status)
			}
		}
	}
}

/////////////////// ELEVATOR CLASS ///////////////////////

type Elevator struct {
	ID                    int
	status                string
	currentFloor          int
	direction             string
	Door                  Door
	floorRequestList      []int
	completedRequestsList []int
	amountOfFloors        int
}

func NewElevator(_id int, _amountOfFloors int) Elevator {
	return Elevator{
		ID:                    _id,
		status:                "idle",
		currentFloor:          0,
		direction:             "idle",
		Door:                  Door{ID: 0},
		floorRequestList:      []int{},
		completedRequestsList: []int{},
		amountOfFloors:        _amountOfFloors,
	}
}

/////////////////// DOOR CLASS ///////////////////////

type Door struct {
	ID     int
	status string
}

func NewDoor(_id int) Door {
	return Door{
		ID:     _id,
		status: "opened",
	}
}

/////////////////// CallButton CLASS ///////////////////////

type CallButton struct {
	ID        int
	status    string
	floor     int
	direction string
}

func NewCallButton(_id int, _floor int, _direction string) CallButton {
	return CallButton{
		ID:        _id,
		status:    "idle",
		floor:     _floor,
		direction: _direction,
	}
}

/////////////////// FloorRequestButtons CLASS ///////////////////////

type FloorRequestButtons struct {
	ID     int
	status string
	floor  int
}

func NewFloorRequestButtons(_id int, _floor int) FloorRequestButtons {
	return FloorRequestButtons{
		ID:     _id,
		status: "idle",
		floor:  _floor,
	}
}

func main() {
	battery1 := Battery{ID: 1, amountOfColumns: 4, amountOfFloors: 60, amountOfElevatorPerColumn: 5, amountOfBasements: 6}
	battery1.assignElevator(8, "up", 0, "down")

	/* fmt.Printf("%+v", value) */
	/* fmt.Printf("%+v", battery1.columnsList[1]) */

}
