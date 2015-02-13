package Driver // where "driver" is the folder that contains io.go, io.c, io.h, channels.go, channels.c and driver.go

/*
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
*/

import  "C"
		."time"
		."math"

//Constants

const N_FLOORS = 4 
const N_BUTTONS = 3

// Variables

var lastDir = 0
var lampMatrix = [N_FLOORS][N_BUTTONS] int{
	{LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
	{LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
	{LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
	{LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
}

var buttonMatrix = [N_FLOORS][N_BUTTONS] int {
	{FLOOR_UP1, FLOOR_DOWN1, FLOOR_COMMAND1},
	{FLOOR_UP2, FLOOR_DOWN2, FLOOR_COMMAND2},
	{FLOOR_UP3, FLOOR_DOWN3, FLOOR_COMMAND3},
	{FLOOR_UP4, FLOOR_DOWN4, FLOOR_COMMAND4},
}


func ElevStartUp() {
	//Initialization og hardware
	if(IoInit() != nil){return 0}

	//Turing off all lamps
	for i := 0; i < N_FLOORS; i++{
		if(i != 0){SetButtonLamp(BUTTON_CALL_DOWN, i, false)
		if(i != N_FLOORS-1){SetButtonLamp(BUTTON_CALL_UP, i, false)}
		SetButtonLamp(BUTTON_COMMAND, i, false)
	}

	//Turing off 
	SetStopLamp(false)
	SetDoorLamp(false)
	floorLamp()



	return 1
}



// gets current floor
func GetCurrentFloor() int {
	if(IoReadBit(SENSOR1)){return = 0}
	else if (IoReadBit(SENSOR2)){return =1}
	else if (IoReadBit(SENSOR3)){return =2}
	else if (IoReadBit(SENSOR4)){return =3}
	else {return = -1}
}
//
// Get button
// Get stop button
//
// Settin floor lamp    Denne kan vi kjøre i bakgrunnen???
func floorLamp() {

	floor = GetCurrentFloor()

	if (floor != -1){
		if ((floor & 0x02) != 0){IoSetBit(FLOOR_IND1)}
		else {IoClearBit(FLOOR_IND1)}
		if ((floor & 0x01) != 0) {IoSetBit(FLOOR_IND2)}
		else {IoClearBit(FLOOR_IND2)}
	}
	//go floorLamp()
}

func readButtons() {
	for i :=0 i < N_FLOORS; i++){
		for j := 0; j < N_BUTTONS; j++ {
			if(IoReadBit(buttonMatrix[i][j])){SetButtonLamp(i,j,true)}
		}
}
	
}
// Setting button lamp
func SetButtonLamp(floor int, button int, on bool) {  //// ON er et dårlig navn!!!
	if (on){IoSetBit(lampMatrix[floor][button])}
	else {IoClearBit(lampMatrix[floor][button])}
}

//Setting stop lamp
func SetStopLamp(stop bool) {
	if(stop){IoSetBit(LIGHT_STOP)}
	else {IoClearBit(LIGHT_STOP)}
}

//Setting door lamp
func SetDoorLamp(door bool) {
	if(door){IoSetBit(DOOR_OPEN)}
	else {IoClearBit(DOOR_OPEN)}
}



// may need "speed variable" to stop det elevator or use "dir"
func SetMotorDir(dir int) {

	//Stopping the elevator
	if(dir == 0){
		if(lastDir == 1){IoSetBit(MOTORDIR)}
		else if(lastDir == -1){IoClearBit(MOTORDIR)}
		Sleep(10*Millisecond) // usikker på om millisecond er en variable
	}
	//Starting the elevator going up
	else if(dir == 1){IoClearBit(MOTORDIR)}
	//Starting the elevator going down
	else if(dir == -1){IoSetBit(MOTORDIR)}

	lastDir = dir
	//Writing new speed to motor
	speed = Abs(dir*(300)) //ANNER IKKE HVA SOM ER NORMAL FART!!!
	IoWriteAnalog(MOTOR, int(2048+4*float64(speed)))
	
}
